package counter

import (
	"sync/atomic"
)

type Counter struct {
	count uint64
	max   uint64
}

func (c *Counter) Inc() uint64 {
	current := atomic.AddUint64(&c.count, 1)
	c.UpdateMax(current)
	return current
}

func (c *Counter) Dec() uint64 {
	return atomic.AddUint64(&c.count, ^uint64(0))
}

func (c *Counter) Reset() {
	atomic.StoreUint64(&c.count, uint64(0))
}

func (c *Counter) Counter() uint64 {
	return atomic.LoadUint64(&c.count)
}

func (c *Counter) UpdateMax(current uint64) bool {
	return atomic.CompareAndSwapUint64(&c.max, current-1, current)
}

func (c *Counter) Max() uint64 {
	return atomic.LoadUint64(&c.max)
}
