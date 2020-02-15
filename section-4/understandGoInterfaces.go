package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Shuffler interface {
	Len() int
	Swap(i, j int)
}

type StackInt []int

func (s StackInt) Len() int {
	return len(s)
}
func (s StackInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type StackString []string

func (s StackString) Len() int {
	return len(s)
}
func (s StackString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Randomise(s Shuffler) {
	fmt.Printf("The sequence before: %v\n", s)
	rand.Shuffle(s.Len(), s.Swap)
	fmt.Printf("The sequence after: %v\n", s)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	Randomise(StackInt{1, 2, 3, 4, 5})
	Randomise(StackString{"the", "quick", "brown", "fox"})

}
