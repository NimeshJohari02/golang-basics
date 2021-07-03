package main

import "fmt"
//NOTE package level means global scope

var  x int
var y string
var z bool

func main() {
	fmt.Printf("%T\t%v",x,x)
	fmt.Printf("%T\t%v",y,y)
	fmt.Printf("%T\t%v",z,z)
// %v helps us print the default values assigned to them also known as the Zero Value
}