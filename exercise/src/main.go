package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"./poetry"
)

var cache = make(map[string]poetry.Poetry)

type config struct {
	Port  int    `json:"port"`
	Route string `json:"route"`
}

func poemsRoute(w http.ResponseWriter, r *http.Request) {
	var filename = r.URL.Query().Get("filename")
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid query parameter\n")
		return
	}

	f, ok := cache[filename]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: File %q not found\n", filename)
		return
	}

	// since a poem implements the Stringer interface, it calls the String() method
	fmt.Fprintf(w, "%s\n", f)
}

func loadConfig() (config, error) {
	var c config
	var configName string

	// Sets default command line arguments and loads them
	flag.StringVar(&configName, "config", "config.json", "Used to load config parameters of the server")
	flag.Parse()

	f, e := os.Open(configName)
	if e != nil {
		return c, e
	}
	defer f.Close()

	// populates the config
	json.NewDecoder(f).Decode(&c)
	return c, nil
}

type customPoem struct {
	name string
	p    poetry.Poetry
	e    error
}

func worker(filePathChan chan string, poemChan chan customPoem, i int) {
	filePath := <-filePathChan
	p, e := poetry.LoadPoem(filePath)
	var fileName = strings.ReplaceAll(filePath, "public/", "")
	poemChan <- customPoem{fileName, p, e}
}

func loadBalancer(filePath string, filePathChan chan string) {
	filePathChan <- filePath
}

func loadFiles() {
	rd, e := ioutil.ReadDir("public")
	if e != nil {
		log.Fatalf("Error: %s\n", e)
	}

	var nRd = len(rd)
	var filePathChan = make(chan string)
	var peomChan = make(chan customPoem)
	for i := 0; i < nRd; i++ {
		go worker(filePathChan, peomChan, i)
	}

	var nExpectedFiles = 0
	for _, df := range rd {
		if df.IsDir() {
			continue
		}

		nExpectedFiles += 1
		var fileName = df.Name()
		go loadBalancer(fmt.Sprintf("public/%s", fileName), filePathChan)
	}

	for i := 0; i < nExpectedFiles; i++ {
		select {
		case p := <-peomChan:
			if p.e != nil {
				log.Fatalf("Error: %s", p.e)
			}
			cache[p.name] = p.p
		}
	}
}

func init() {
	loadFiles()
}

func main() {
	var c config
	var e error
	if c, e = loadConfig(); e != nil {
		log.Fatalf("Error: %s\n", e)
	}

	http.HandleFunc(c.Route, poemsRoute)

	var port = fmt.Sprintf(":%d", c.Port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
