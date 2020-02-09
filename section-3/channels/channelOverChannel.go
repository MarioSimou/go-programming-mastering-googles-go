package main

import (
	"fmt"
	"time"
)

func emit(chanChannel chan chan string, done chan bool) {
	var wordsChannel = make(chan string)
	chanChannel <- wordsChannel
	var i = 0
	var words = []string{"the", "quick", "brown", "fox"}
	var timer = time.NewTimer(time.Second * 4)
	for {
		select {
		case wordsChannel <- words[i]:
			i++
			if i >= len(words) {
				i = 0
			}
		case <-timer.C:
			fmt.Println("time hits..")
			// close(wordsChannel)
			done <- true
			close(done)
			close(wordsChannel)
			return
		}
	}

}

func main() {
	var chanChannel = make(chan chan string)
	var done = make(chan bool)

	go emit(chanChannel, done)
	wordsChannel := <-chanChannel

daemon:
	for {
		select {
		case word := <-wordsChannel:
			fmt.Printf("Word: %s\n", word)
		case <-done:
			break daemon
		}
	}
	fmt.Println("FINISH PROGRAM")
}
