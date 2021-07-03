package main

import (
	"fmt"
)
type person struct{
	fname string
	lname string
}
type Agent struct {
	person
	secretId int
}
func (p person) myfunc(){
fmt.Println("I was Called By Person with Name " ,p.fname)
}
func (sa Agent) myfunc(){
fmt.Println("I was Called By Agent with Name " ,sa.fname)
}
type human interface{
	myfunc()
}
func DemoFunc(h human ){
	switch h.(type){
	case person:
		//What we have done below is told the compiler that the h being passed is of type person 
		//s so use the fname associated with person type and do not create conflict 
		fmt.Println("I was Called inside demo by human also person" , h.(person).fname)
	case Agent:
		fmt.Println("I was Called inside Demo by Human Also an Agent ", h.(Agent).fname )
	}
}
func main() {
	p:=person{
		fname: "Nimesh",
		lname: "johari",
	}	
	agent := Agent{
		person: person{
			fname: "Shruti",
			lname: "Johari",
		},
		secretId: 12312,
	}
	DemoFunc(p);
	DemoFunc(agent)
}