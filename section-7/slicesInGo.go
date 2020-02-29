package main

import "fmt"

func main(){
	var all = []int{1,2,3,4,5}
	var copy = all[1:3]

	copy[0] = 10 // all[1] and copy[0] change to 10
	all[2] = 2 * all[2] // a[[2]] and copy[1] change to 6

	// the values 30,31 will replace 4,5 in all array (until it reaches to maximum capacity (5))
	// the values 30,31,32,33 will be added to copy, with its capacity to switch to 8
	copy = append(copy,30)
	copy = append(copy,31)
	copy = append(copy,32)
	copy = append(copy,33)

	fmt.Printf("all: %v\tcap all: %d\n", all, cap(all))
	fmt.Printf("copy:%v\tcap all: %d\n", copy, cap(copy))
}