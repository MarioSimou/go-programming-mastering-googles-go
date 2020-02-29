package main

import "fmt"

type Person struct {
	Name string
}

func main(){
	var p1 = Person{"John"}
	var p2 = Person{"Paul"}
	var set = make(map[Person]struct{})

	set[p1] = struct{}{}
	set[p2] = struct{}{}

	if _, ok := set[p1]; ok {
		fmt.Printf("Match\n")
	}else {
		fmt.Printf("No Match\n")
	}
}