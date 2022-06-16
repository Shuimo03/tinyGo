package mutex

import "sync"

type Counter struct {
	CounterType int
	Name        string
	mu          sync.Mutex
	cnt         uint64
}

func (counter *Counter) Incr() {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.cnt++
}

func (counter *Counter) Count() uint64 {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	return counter.cnt
}
