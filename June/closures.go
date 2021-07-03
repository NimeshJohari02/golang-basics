package main

import "fmt"
func main() {
	// basically closures are inline functions in js
	add := func(x,y int)int{
		return x+y
	}
	x:=0;
	increment := func () int {
		x++;
		return x;
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(add(2,32))
}