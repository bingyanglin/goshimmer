package prng_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/iotaledger/goshimmer/packages/prng"
)

func TestResolveNextTimePoint(t *testing.T) {
	assert.EqualValues(t, 105, prng.ResolveNextTimePointSec(103, 5))
	assert.EqualValues(t, 110, prng.ResolveNextTimePointSec(105, 5))
	assert.EqualValues(t, 105, prng.ResolveNextTimePointSec(100, 5))
	assert.EqualValues(t, 100, prng.ResolveNextTimePointSec(97, 5))
}

func TestUnixTsPrng(t *testing.T) {
	unixTsRng := prng.NewUnixTimestampPRNG(1 * time.Second)
	unixTsRng.Start()
	defer unixTsRng.Stop()

	var last float64
	for i := 0; i < 3; i++ {
		r := <-unixTsRng.C()
		assert.Less(t, r, 1.0)
		assert.Greater(t, r, 0.0)
		assert.NotEqual(t, last, r)
		last = r
	}
}
