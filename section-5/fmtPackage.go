package main

import (
	"fmt"
)

type myString struct {
	content string
}

// implements the Stringer interfaces under fmt.Stringer
func (ms myString) String() string {
	return "hello world"
}

func main() {
	var s = myString{"some string\n"}
	fmt.Printf("% v\n", s)
}
