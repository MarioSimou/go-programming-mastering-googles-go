package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func poemHandler(w http.ResponseWriter, r *http.Request) {
	// read file and serve it
	var fileName = r.URL.Query().Get("file")
	http.ServeFile(w, r, fileName)
}

type config struct {
	Port  string `json:"port"`
	Route string `json:"route"`
}

func main() {
	var start = time.Now()
	fmt.Printf("Starting at: %s\n", start.Format(time.UnixDate))
	var c config
	var configFileName = flag.String("conf", "config.json", "Default config file")
	flag.Parse() // don't forget to call the parse method

	f, e := ioutil.ReadFile(*configFileName)
	if e != nil {
		log.Fatalf("Error: %s \n", e)
	}
	json.NewDecoder(bytes.NewReader(f)).Decode(&c)

	http.HandleFunc(c.Route, poemHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))

	var port = fmt.Sprintf(":%s", c.Port)
	var elapsed = time.Now().Sub(start)

	fmt.Printf("Elapsed time: %s\n", elapsed)
	log.Fatalln(http.ListenAndServe(port, nil))
}
