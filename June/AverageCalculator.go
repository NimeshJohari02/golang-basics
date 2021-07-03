package main

import "fmt"
func main(){
	var arr[5]float64
	var total float64=0;
	for i:=62;i<67;i++{
		arr[i-62]=float64(i)
		total+=float64(i)
	}
	fmt.Print(total/float64(len(arr)));
}