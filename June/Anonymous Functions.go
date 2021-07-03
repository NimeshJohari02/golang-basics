package main

import "fmt"
func root() func(a int) int{
	return func(a int) int {
		return a*a;
	}
}
func main() {
	fmt.Println("Anonymous Function Demo")
	func(){
		fmt.Println("Anonymous Called ")
	}()
	func(a int , b int ){
		fmt.Println(a+b);
	}(23,24)
	// This is a functional Expression Given Below 
	f1:=func(a int , b int){fmt.Println(a+b)}
	f1(23,67);
	rf:=root();
	fmt.Println(rf(23))
}