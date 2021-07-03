package main

import "fmt"

type mytype int

var a int
var b mytype

func main() {
	b = 90
	a = 20
	fmt.Println(b)
	fmt.Println(a)
	/*	a=b
		Note : We Can not do this as b is of type mytype and a is of the type integer thus we can not store the value of b to a
	 What we can rather do is */
	a = int(b)
	fmt.Println(a)

}
