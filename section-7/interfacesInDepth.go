package main

import (
	"fmt"
	"strings"
)

type Yeller interface {
	String() string
	Change(string)
	Blank()
	Len() int
}

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
func (ls *LoudString) Len() int {
	return len(ls.s)
}


func main(){
	var ls = NewLoudString()
	var change = ls.Change
	change("Hello")

	// checks if ls structure impleemnts the Yeller and Stringer interface
	var _ Yeller = &ls // the *LoudString implements the interface and not LoudString
	var _ fmt.Stringer = &ls
}