package main

import (
	"fmt"
)

type Context struct {
	done chan bool
}
func NewContext() *Context {
	var done = make(chan bool)
	return &Context{done}
}

func (c *Context) Stop(){
	close(c.done)
}

func (c *Context) GetDone() chan bool {
	return c.done
}

type Counter struct {
	ctx *Context
	c chan int
	i int
}

func NewCounter() *Counter {
	var counter Counter
	counter.c = make(chan int)
	counter.ctx = NewContext()

	var done = counter.ctx.GetDone()
	go func(){
		for {
			select {
			case counter.c <- counter.i:
				counter.i = counter.i + 1
			case <- done:
				fmt.Printf("Counter terminated!\n")
				return
			}
		}
	}()

	return &counter
}

func (c *Counter) GetSource() <- chan int {
	return c.c
}

func (c *Counter) Stop(){
	c.ctx.done <- true
}


func main(){
	var counter = NewCounter()
	var read = counter.GetSource()

	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	counter.Stop()
}