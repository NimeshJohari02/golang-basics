package main

import "fmt"
type person struct{
	fname string
	lname string
}
func (personPtr *person) SayHello(){
	fmt.Println("This is Person Who Says Hello By Reciving personPtr")
}
type human interface{
	SayHello() 
}
func getInformation(humanObject human){
	fmt.Println("GetInformation" )
 humanObject.SayHello()
}
func main() {
p:=person{
	fname: "Nimesh",
	lname: "Johari",
}
p.SayHello()
getInformation(&p)
}