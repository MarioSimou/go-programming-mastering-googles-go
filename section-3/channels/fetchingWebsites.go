package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getPage(url string) (int, error) {
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

func getter(url string, sizeChan chan string, errorChan chan error) {
	size, e := getPage(url)
	if e != nil {
		errorChan <- e
	}
	sizeChan <- fmt.Sprintf("The %s has length of %d", url, size)
}

func main() {
	var webpages = []string{
		"https://www.google.com",
		"https://www.udemy.com",
		"https://www.facebook.com",
	}

	var sizeChan = make(chan string)
	var errorChan = make(chan error)
	for _, website := range webpages {
		go getter(website, sizeChan, errorChan)
	}

	for i := 0; i < len(webpages); i++ {
		select {
		case size := <-sizeChan:
			fmt.Println(size)
		case e := <-errorChan:
			log.Fatalln("Error: ", e)
		}
	}
	fmt.Println("FINISH PROGRAM")

}
