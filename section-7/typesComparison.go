package main

import "fmt"

type MyString struct {
	Value string
}

type MyInterface interface {
	String() string
}


type MyString1 string
func (ms MyString1) String() string {
	return string(ms)
}

type MyString2 string
func (ms MyString2) String() string {
	return string(ms)
}


func main(){
	// fmt.Printf("%t\n", MyString{"Value"} == MyString{"Value"}) // true 
	// fmt.Printf("%t\n", MyString{"Value"} == MyString{"Other Value"}) // false

	var v1 MyString1 = "Hello"
	var v2 MyString2 = "Hello"
	fmt.Printf("%t\n", MyInterface(v1) == MyInterface(v2) ) // false

	fmt.Printf("%t\n", MyInterface(v1) == MyInterface(v1)) // true

	var m1 = make(map[string]string)
	m1["Value"] = "Value"
	var m2 = make(map[string]string)
	m2["Value"] = "Value"
	// fmt.Printf("%t\n", m1 == m2) // error (map can only be compared to nil)
	
}