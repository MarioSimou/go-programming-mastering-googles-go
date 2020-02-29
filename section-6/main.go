package main

import (
	"fmt"
)

type error interface {
	Error() string
}

type custom struct{}

func (c custom) Error() string {
	return "some error"
}

func main() {
	var c custom
	fmt.Printf("Error: %s\n", c)

}
