package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type webpage struct {
	url  string
	body []byte
	err  error
}

func (w *webpage) get() {
	res, e := http.Get(w.url)
	if e != nil {
		w.err = e
	}
	defer res.Body.Close()

	w.body, e = ioutil.ReadAll(res.Body)
	if e != nil {
		w.err = e
	}
}

func (w *webpage) isOk() bool {
	return w.err == nil
}

func main() {
	var w = &webpage{"https://www.google.com", nil, nil}
	w.get()
	if !w.isOk() {
		log.Fatalln(w.err)
	}
	fmt.Printf("Called webpage %s with a size of %d\n", w.url, len(w.body))
}
