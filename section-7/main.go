package main

import (
	"flag"
	"fmt"
	"time"
)

func main(){
	var defaultDuration = 5 * time.Second * time.Nanosecond // ns
	var timeout = flag.Duration("timeout", defaultDuration, "api timeout")
	flag.Parse()

	fmt.Printf("%d\n", *timeout)
}