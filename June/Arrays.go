package main

import (
	"fmt"
)

func main(){
	var arr[5] int;
	fmt.Println("Arr Initally = ",arr);
	for i:=0;i<5;i++{
		arr[i]=i;
	}
	fmt.Println(" Length ",len(arr));
	fmt.Println("Arr Finally = ",arr);
	arrInit:=[5] int{1,4,9,16,25};
	fmt.Println("Arr Initally = ",arrInit);
	//2d arratys
	var twoDimensionArray [3][3]int
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			twoDimensionArray[i][j]=i+j
		}
	}
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			fmt.Print(twoDimensionArray[i][j],"   ")
		}
		fmt.Print("\n")
	}
}