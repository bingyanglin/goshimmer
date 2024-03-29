package curl

import (
	"fmt"
	"strconv"
	"time"

	"github.com/iotaledger/goshimmer/packages/ternary"
)

type HashRequest struct {
	input  ternary.Trits
	output chan ternary.Trits
}

type BatchHasher struct {
	hashRequests chan HashRequest
	tasks        chan []HashRequest
	hashLength   int
	rounds       int
}

func NewBatchHasher(hashLength int, rounds int) *BatchHasher {
	this := &BatchHasher{
		hashLength:   hashLength,
		rounds:       rounds,
		hashRequests: make(chan HashRequest),
		tasks:        make(chan []HashRequest, NUMBER_OF_WORKERS),
	}

	go this.startDispatcher()
	this.startWorkers()

	return this
}

func (this *BatchHasher) startWorkers() {
	for i := 0; i < NUMBER_OF_WORKERS; i++ {
		go func() {
			for {
				this.processHashes(<-this.tasks)
			}
		}()
	}
}

func (this *BatchHasher) startDispatcher() {
	for {
		collectedHashRequests := make([]HashRequest, 0)

		// wait for first request to start processing at all
		collectedHashRequests = append(collectedHashRequests, <-this.hashRequests)

		// collect additional requests that arrive within the timeout
	CollectAdditionalRequests:
		for {
			select {
			case hashRequest := <-this.hashRequests:
				collectedHashRequests = append(collectedHashRequests, hashRequest)

				if len(collectedHashRequests) == strconv.IntSize {
					break CollectAdditionalRequests
				}
			case <-time.After(50 * time.Millisecond):
				break CollectAdditionalRequests
			}
		}

		this.tasks <- collectedHashRequests
	}
}

func (this *BatchHasher) processHashes(collectedHashRequests []HashRequest) {
	if len(collectedHashRequests) > 1 {
		// multiplex the requests
		multiplexer := ternary.NewBCTernaryMultiplexer()
		for _, hashRequest := range collectedHashRequests {
			multiplexer.Add(hashRequest.input)
		}
		bcTrits, err := multiplexer.Extract()
		if err != nil {
			fmt.Println(err)
		}

		// calculate the hash
		bctCurl := NewBCTCurl(this.hashLength, this.rounds)
		bctCurl.Reset()
		bctCurl.Absorb(bcTrits)

		// extract the results from the demultiplexer
		demux := ternary.NewBCTernaryDemultiplexer(bctCurl.Squeeze(243))
		for i, hashRequest := range collectedHashRequests {
			hashRequest.output <- demux.Get(i)
			close(hashRequest.output)
		}
	} else {
		var resp = make(ternary.Trits, this.hashLength)

		curl := NewCurl(this.hashLength, this.rounds)
		curl.Absorb(collectedHashRequests[0].input, 0, len(collectedHashRequests[0].input))
		curl.Squeeze(resp, 0, this.hashLength)

		collectedHashRequests[0].output <- resp
		close(collectedHashRequests[0].output)
	}
}

func (this *BatchHasher) Hash(trits ternary.Trits) chan ternary.Trits {
	hashRequest := HashRequest{
		input:  trits,
		output: make(chan ternary.Trits, 1),
	}

	this.hashRequests <- hashRequest

	return hashRequest.output
}

const (
	NUMBER_OF_WORKERS = 1000
)
