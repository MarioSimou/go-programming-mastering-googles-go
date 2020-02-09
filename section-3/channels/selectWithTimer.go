package main

import (
	"fmt"
	"time"
)

// function that sends messages for X duration and then closes the available channels
func emit(wordsChan chan string, doneChan chan bool) {
	var words = []string{"the", "quick", "brown", "fox"}
	var i = 0
	var timer = time.NewTimer(time.Second * 5)
	for {
		select {
		case wordsChan <- words[i]:
			i++
			if i >= len(words) {
				i = 0
			}
		case <-timer.C:
			close(wordsChan)
			close(doneChan)
			return
		}
	}
}

func main() {
	var wordsChan = make(chan string)
	var doneChan = make(chan bool)

	go emit(wordsChan, doneChan)

daemon:
	for {
		select {
		case word := <-wordsChan:
			fmt.Printf("Word: %s\n", word)
		case <-doneChan:
			break daemon
		}
	}

	fmt.Printf("FINISH PROGRAM\n")
}
