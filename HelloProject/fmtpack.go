package main

import "fmt"

var x = 43

func main() {
	fmt.Println(x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%o\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\n", x)
	s := fmt.Sprintf("%b\t%o\t%#X\t", x, x, x)
	fmt.Println(s)

}
