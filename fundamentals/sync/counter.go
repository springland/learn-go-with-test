package counter

import "sync"
type Counter struct{
	 val int

	 mu sync.Mutex
}

func (c *Counter) Inc(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val ++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}