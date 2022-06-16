package waitGroup

import (
	"sync"
	"time"
)

//Counter 线程安全的计数器
type Counter struct {
	mu    sync.Mutex
	count uint64
}

// Incr 对计数值+1
func (counter *Counter) Incr() {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.count++
}

//CurrCount currCount获取当前计数
func (counter *Counter) CurrCount() uint64 {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	return counter.count
}

// Worker sleep 1秒，然后计数值加1
func Worker(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	counter.Incr()
}
