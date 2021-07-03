package main

import "fmt"

type chicken int

var x chicken
var y int

func main() {
	fmt.Printf("%v\t%T\t", x, x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\t",y)

}
