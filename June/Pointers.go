package main

//! Pointers can also be created usign new keywods

import "fmt"
func ReduceQuantity(xPtr *int){
	*xPtr-=2;
}
func setX(intPtr *int){
	*intPtr = 20
}
func main() {
	x:=14;
	xPtr:= new(int);
	setX(xPtr);
	fmt.Println(*xPtr)	
	ReduceQuantity(&x)
	fmt.Println(x)	
}