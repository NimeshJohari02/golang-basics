package main

import "fmt"
type person struct{
	fname string
	lname string
	address string
	over18 bool
}
func main() {
	// Make A Value Of Type Person 
	person1:=person{
		fname:"Nimesh",
		lname: "Johari",
		address: "Lucknow-UP",
		over18: true,
	}
	fmt.Println(person1)
	if person1.over18{
		fmt.Println("Legal To Drive")
	}
}