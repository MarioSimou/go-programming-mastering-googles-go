package main

import (
	"fmt"
)

func emit(wordsChan chan string, doneChan chan bool) {
	var words = []string{"the", "quic", "brown", "fox"}
	var i = 0
	for {
		select {
		case wordsChan <- words[i]:
			i++
			if i >= len(words) {
				i = 0
			}
		case <-doneChan:
			fmt.Printf("CLOSING CHANNEL\n")
			// close(doneChan)
			doneChan <- true
			return
		}
	}
}

func main() {
	var wordsChan = make(chan string)
	var doneChan = make(chan bool)

	go emit(wordsChan, doneChan)

	for i := 0; i < 100; i++ {
		fmt.Printf("i: %d\tWord: %s\n", i, <-wordsChan)
	}
	doneChan <- true
	<-doneChan // waits until it receives a value
	fmt.Printf("Program terminates...")
}
