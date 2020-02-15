package shuffler

import "math/rand"

type Shuffler interface {
	Len() int
	Swap(i, j int)
}

func Shuffle(s Shuffler) Shuffler {
	rand.Shuffle(s.Len(), s.Swap)
	return s
}
