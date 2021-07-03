package main

import "fmt"

type chicken int
var x chicken
func main() {
	fmt.Printf("%v\t%T\t",x,x)
	x=42
	fmt.Println(x)
}
