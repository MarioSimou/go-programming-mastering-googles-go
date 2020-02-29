package main

import (
	"flag"
	"fmt"

	s "./shouter"
)


func main(){
	var version = flag.String("version", "1.0.0", "Get app version")
	flag.Parse()

	if v := *version; v != "" {
		fmt.Printf("Version: %s\n", v)
	}

	var str = "my name is John"
	fmt.Printf("%s\n", s.Shout(str))
}