package main

import "fmt"

func whatIsThis(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("it is int with a value of %v\n", v)
	case string:
		fmt.Printf("it is string with a value of %v\n", v)
	default:
		fmt.Println("other")
	}
}

func main() {
	whatIsThis("hello")
}
