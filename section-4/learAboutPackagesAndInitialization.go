package main

import (
	"fmt"

	s "./shuffler"
)

type stackInt []int

func (si stackInt) Len() int {
	return len(si)
}
func (si stackInt) Swap(i, j int) {
	si[i], si[j] = si[j], si[i]
}

func main() {
	var sInt = stackInt{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Printf("Shuffled result: %v", s.Shuffle(sInt))
}
