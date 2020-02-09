package main

import (
	"fmt"
	"math/rand"
	"time"
)

// with nil channels we can terminate the execution of a channel
func reader(numChan chan int) {
	var terminator = time.NewTimer(time.Second * 10)
	for {
		select {
		case num := <-numChan:
			fmt.Printf("Number: %d\n", num)
		case <-terminator.C:
			numChan = nil
		}
	}
}

func writer(numChan chan int) {
	var terminator = time.NewTimer(time.Second * 2)
	var restarted = time.NewTimer(time.Second * 5)
	ch := numChan
	for {
		select {
		case numChan <- rand.Intn(42):
		case <-terminator.C:
			fmt.Println("TERMINATES FOR 3 SECONDS UNTIL IT RESTARTS AGAIN")
			numChan = nil
		case <-restarted.C:
			numChan = ch
		}
	}
}

func main() {
	var numChan = make(chan int)
	go reader(numChan)
	go writer(numChan)

	time.Sleep(time.Second * 10)
	fmt.Println("PROGRAM TERMINATES")
}
