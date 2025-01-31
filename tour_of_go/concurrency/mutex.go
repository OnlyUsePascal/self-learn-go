package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct{
	mu sync.Mutex
	cnt int
}

func (c *Counter) counterInc(){
	// try commenting this :)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cnt++
}

func (c *Counter) counterGet() int{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cnt
}

func Mutex_basic(){
	c := Counter{cnt : 0}
	
	for i := 0; i < 1000; i++ {
		go c.counterInc()
	}
	
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(c.counterGet())
}