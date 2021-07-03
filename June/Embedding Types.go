package main

import (
	"fmt"
)
type person struct {
 name string
 lanme string 
}
type employee struct{
	person
	company string
	city string

}
func main() {
	p1:=employee{
		person: person{
			name: "Nimesh",
			lanme: "Johari",
		},
		company: "Unemployed",
		city:`Lucknow`,
	}	
	// ! Once you run this you will see that you 
	// dont need to use p1.person.name you can use that but it is automatically Promoted to the uppper level 
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.lanme)
	dummy:=struct{
 name string
 lname string 
	}{
		name: "Shruti",
		lname: "Johari",
	}
fmt.Println(dummy.name)	
fmt.Println(dummy.lname)	


}
