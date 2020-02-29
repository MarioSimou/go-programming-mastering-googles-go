package main

import (
	"fmt"
	"strings"
)

type LoudString struct {
	s string
}

func NewLoudString() LoudString {
	return LoudString{}
} 

func (ls *LoudString) String() string {
	return ls.s
}

func (ls *LoudString) Change(s string){
	ls.s = strings.ToUpper(s)
}

func (ls *LoudString) Blank(){
	ls.s = ""
}


func main(){
	var ls = NewLoudString()
	ls.Change("Hello")
	fmt.Printf("Value: %s\n", ls.String())

	// creating a closure
	var change = ls.Change
	change("Hello Hello")
	fmt.Printf("Value: %s\n", ls.String())

}