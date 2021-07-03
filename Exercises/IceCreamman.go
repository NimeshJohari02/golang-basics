package main

import "fmt"
type person struct{
	fname string
	lname string
	favourites []string
	// myMap map[string][]string
}
func main() {	
	person1:=person{
		fname: "Nimesh",
		lname: "Johari",
		favourites: []string{"Chocolate","Vannila"},
	}
	for _,v:=range person1.favourites{
		fmt.Println("Favourites = ",v)
	}
}