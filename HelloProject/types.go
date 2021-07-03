package main

import "fmt"
// declaring a variable with the IDENTIFIER "z" of TYPE = integer
var z =12
var x ="It is What it is "
func main() {
	fmt.Print("Welcome to Types in Go")
	fmt.Println(x)
	// To print type WE are usiing the printf not println
	fmt.Printf("%T \n ",x)
	fmt.Println(z)
	fmt.Printf("%T \n",z)
	/* We can not set x = 47 or something because data type pehle set kar diya hai */
}