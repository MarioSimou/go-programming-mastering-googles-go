package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main(){
	var m = make(map[io.Writer]bool)
	m[os.Stdout] = true // Stdout and Stderr implement the Writer interface
	m[os.Stderr] = false // 
	var bf = bytes.NewBuffer([]byte{}) // buffer implements the Writer Interface
	m[bf] = true


	for w, enabled := range m {
		if enabled {
			fmt.Fprintf(w,"Hello world\n")
		}
	}

	fmt.Printf("Buffer size: %d\n", bf.Len())
}