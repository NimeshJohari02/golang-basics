package main

import "fmt"
func main() {
	x:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14}
	fmt.Println(x)
	fmt.Println(sum(x...))
	fmt.Println(odd(sum,x...))
	fmt.Println(EvenSum(sum,x...))
}
func sum(arr ...int) int{
	total:=0;
	for _, v := range arr {
	total+=v	
	}
	return total
}
func EvenSum(f func(arr ...int) int , evenNumberArr ...int) int {
	var dummyList []int 
	for _, v := range evenNumberArr {
		if(v%2==0){
			dummyList=append(dummyList, v);
		}
	}
	return f(dummyList...)
}
func odd( f func(arr ...int) int , oddArr...int) int {
	var dummyList []int
	for _, v := range oddArr {
		if(v%2!=0){
			dummyList=append(dummyList, v);
		}	
	}
	return f(dummyList...)
}