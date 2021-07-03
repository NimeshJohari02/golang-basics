package main

import "fmt"
func main() {
	x:=[]int{32,42,23,12,21,22,31,53}
	for i,v:=range x{
		fmt.Println(i,v)
	}
	fmt.Printf("%T",x);
	sliceDemo:=[]int{42,43,44,45,46,47,48,49,50,50,51}
	sliceDemo = append(sliceDemo, 52)
	fmt.Println(sliceDemo)
	sliceDemo = append(sliceDemo, 53,54,55)
	sliceDemo=append(sliceDemo, []int{56,57,58,59}...)
	sliceDemo=append(sliceDemo[:4], sliceDemo[7:]...)
	TwoDimSlice:=[][]string{
		{"This is row 1 Col 1","Row 1 Col2"},
		{"This is row 2 Col 1","Row 2 Col2"},
		{"This is row 3 Col 1","Row 3 Col2"},
	}
	for _,v:=range TwoDimSlice{
		for _,innerVal := range v{
			fmt.Println(innerVal)
		}
	}
	 fmt.Println(sliceDemo)
}