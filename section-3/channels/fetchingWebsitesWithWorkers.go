package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchPage(url string) (int, error) {
	r, e := http.Get(url)
	if e != nil {
		return 0, e
	}
	defer r.Body.Close()

	bf, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return 0, e
	}
	return len(bf), nil
}

func worker(urlsChan chan string, sizeChan chan string, i int) {
	url := <-urlsChan
	s, e := fetchPage(url)
	if e != nil {
		sizeChan <- fmt.Sprintf("Error: %s\n", e.Error())
	} else {
		sizeChan <- fmt.Sprintf("The size of %s is %d (%d)\n", url, s, i)
	}
}

func loadBalancer(urlsChan chan string, url string) {
	urlsChan <- url
}

func main() {
	var websites = []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.gmail.com",
	}
	var sizeChan = make(chan string)
	var urlsChan = make(chan string)

	// initialize 10 workers
	for i := 0; i < 10; i++ {
		go worker(urlsChan, sizeChan, i)
	}

	for _, website := range websites {
		go loadBalancer(urlsChan, website)
	}

	for i := 0; i < len(websites); i++ {
		select {
		case size := <-sizeChan:
			fmt.Printf(size)
		}
	}
	fmt.Println("FINISH PROGRAM")
}
