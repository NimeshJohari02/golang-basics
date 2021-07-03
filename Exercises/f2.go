package main

import "fmt"
func twoReturn() (int , string){
	return 34,"This is Nimesh";
}
func f1() string{
	return "Hello Worlds"
}
func sum (x ...int) int {
	total:=0;
	for _, v := range x {
		total+=v
	}
	return total
}
func main() {
	x,y:=twoReturn()	
	fmt.Println(x,y)
	str :=f1();
	fmt.Println(str)
	arr:=[]int{12,413,121,32,12,2312,231,12};
	fmt.Println(sum(arr...))
}