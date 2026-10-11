package syncpkg

import "sync"

type Counter struct {
	mu  sync.Mutex
	val int
}

func (counter *Counter) Value() int {
	return counter.val
}

func (counter *Counter) Inc() {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.val++
}
