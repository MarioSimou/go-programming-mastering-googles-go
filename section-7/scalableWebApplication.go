package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Response struct {
	e error
	url string
	from int
	message string
	duration time.Duration
}

func worker(urlCh chan string, respCh chan Response, i int){
	var start = time.Now()
	var url = <- urlCh
	var resp, e = http.Get(url)
	if e != nil {
		respCh <- Response{e:e, url: url, from: i, duration: time.Now().Sub(start)}		
		return
	}
	if code := resp.StatusCode; code != 200 {
		respCh <- Response{e: e, url: url, from: i, duration: time.Now().Sub(start)}		
		return
	}
	respCh <- Response{message: "ok", url: url, from: i, duration: time.Now().Sub(start)}
}

func writer(respCh chan Response, wg *sync.WaitGroup, nLines int){
	defer wg.Done()
	for i:=0; i < nLines; i++{
		select {
			case response := <- respCh:
				if e := response.e; e != nil {
					fmt.Printf("(%s) (%d): Error (%s): %s\n", response.duration, response.from, response.url, e.Error())
				} else {
					fmt.Printf("(%s) (%d): Ok (%s)\n", response.duration,response.from,response.url)
				}
		}		
	}
}

func ReadStdIn(reader io.Reader, nWorkers int){
	var scanner = bufio.NewScanner(reader)
	var urlCh = make(chan string)
	var respCh = make(chan Response)
	var wg sync.WaitGroup
	var nLines = 0

	fmt.Printf("Creating %d workers....\n", nWorkers)
	for i:=0; i < nWorkers; i++ {
		go worker(urlCh,respCh, i)
	}

	for scanner.Scan() {
		var t = scanner.Text()
		nLines += 1

		if e := scanner.Err(); e != nil {
			log.Fatalf("Reading error: %s\n", e.Error())
		}
		urlCh <- t
	}

	wg.Add(1)
	go writer(respCh, &wg, nLines)
	wg.Wait()
}

func main(){
	var nWorkers = flag.Int("nWorkers", 50, "Number of workers")
	flag.Parse()
	ReadStdIn(os.Stdin, *nWorkers)
}