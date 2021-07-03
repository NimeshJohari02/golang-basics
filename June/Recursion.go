package main

import "fmt"
func getFactorial(x uint) uint{
	if(x==0 || x==1){
		return x
	}
	return x*getFactorial(x-1)
}
func main() {
	fmt.Println(getFactorial(6))
}