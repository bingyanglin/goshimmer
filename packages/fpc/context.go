package fpc

import (
	"sync"
)

// Context defines the context of an FPC instance
type context struct {
	opinionHistory *OpinionMap
	activeTxs      etaMap
	waitingTxs     *txQueue
	parameters     *Parameters
	tick           *tick
}

// NewContext returns a new FPC context
func newContext(p ...*Parameters) *context {
	param := NewParameters()
	if p != nil && p[0] != nil {
		param = p[0]
	}
	return &context{
		opinionHistory: NewOpinionMap(),
		activeTxs:      newEtaMap(),
		waitingTxs:     newTxQueue(),
		parameters:     param,
		tick:           &tick{},
	}
}

type txQueue struct {
	sync.RWMutex
	internal []TxOpinion
}

func newTxQueue() *txQueue {
	return &txQueue{}
}

func (tq *txQueue) Len() int {
	tq.RLock()
	defer tq.RUnlock()
	return len(tq.internal)
}

func (tq *txQueue) Push(txs ...TxOpinion) {
	if txs == nil {
		return
	}
	tq.Lock()
	defer tq.Unlock()
	tq.internal = append(tq.internal, txs...)
}

func (tq *txQueue) Pop(n ...uint) (out []TxOpinion) {
	tq.Lock()
	defer tq.Unlock()

	if n == nil || len(n) == 0 || n[0] > uint(len(tq.internal)) {
		out = make([]TxOpinion, len(tq.internal))
		copy(out, tq.internal)
		tq.internal = nil
		return out
	}
	out = make([]TxOpinion, n[0])
	copy(out, tq.internal[:n[0]])
	tq.internal = tq.internal[n[0]:]
	return out
}

func (c *context) pushTxs(txs ...TxLike) {
	for _, tx := range txs {
		c.opinionHistory.Store(tx.TxHash, tx.ToOpinion())
		c.waitingTxs.Push(tx.ToTxOpinion())
	}
}

// TODO: set max number of txs to add as active
func (c *context) popTxs() {
	newTxs := c.waitingTxs.Pop()
	for _, tx := range newTxs {
		c.activeTxs[tx.TxHash] = &etaResult{
			value: -1,
			count: 0,
		}
	}
}

func (c *context) getActiveTxs() []ID {
	txs := []ID{}
	for tx := range c.activeTxs {
		txs = append(txs, tx)
	}
	return txs
}
