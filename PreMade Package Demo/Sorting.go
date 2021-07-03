package main

import (
	"fmt"
	"sort"
)
type person struct{
	name string
	age int
}
type ByName []person
func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].name < a[j].name }
func main() {
	// x:=[]int{12,31,13,12,43,23,12}
	// fmt.Println(x)	
	// sort.Ints(x);
	// fmt.Println(x)
	x1:=person{
		name: "Nimesh",
		age: 20,
	}
	x2:=person{
		name: "Shruti",
		age: 24,
	}
	x3:=person{
		name: "Prashant",
		age: 39,
	}
	people:=[]person{x1,x2,x3}
	sort.Sort(ByName(people))
	fmt.Println(people)
}