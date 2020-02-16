package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"./poetry"
)

var protectedCache = cache{c: make(map[string]poetry.Poetry)}

type config struct {
	Port  int    `json:"port"`
	Route string `json:"route"`
}

type cache struct {
	sync.Mutex
	e error
	c map[string]poetry.Poetry
}

func poemsRoute(w http.ResponseWriter, r *http.Request) {
	var filename = r.URL.Query().Get("filename")
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid query parameter\n")
		return
	}

	f, ok := protectedCache.c[filename]
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

func main() {
	var c config
	var e error
	if c, e = loadConfig(); e != nil {
		log.Fatalf("Error: %s\n", e)
	}

	rd, e := ioutil.ReadDir("public")
	if e != nil {
		log.Fatalf("Error: %s\n", e)
	}

	var wg sync.WaitGroup
	for _, fd := range rd {
		if fd.IsDir() {
			continue
		}

		wg.Add(1)
		go func(fileName string, wg *sync.WaitGroup) {
			var p, _ = poetry.LoadPoem(fmt.Sprintf("public/%s", fileName))
			protectedCache.Lock()
			protectedCache.c[fileName] = p
			protectedCache.Unlock()

			wg.Done()
		}(fd.Name(), &wg)
	}
	wg.Wait()

	http.HandleFunc(c.Route, poemsRoute)

	var port = fmt.Sprintf(":%d", c.Port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
