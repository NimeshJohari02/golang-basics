package main

import "fmt"
type student struct{
	fname string
	lname string
	marks []int
	average float64
}
// func (r reciever) function_name(parametres) (return(s)){code ... }
//! Note when you attach structure before function name it becomes input for the receiver clause 
//! Basically attach a method to type by placing it as receiver 
func(mystudent student) getAverage() float64 {
	s:=0.0;
	for _,v:=range mystudent.marks{
		s+=float64(v)
	}
	s=float64(s)/float64(len(mystudent.marks))
	return s
}
// Interface says if you have a specified method then you are of that type / interface 
func main() {
	st1:=student{
		fname: "Nimesh",
		lname: "Johari",
		marks: []int{98,97,92,78,69},
	}	
	st1.average=st1.getAverage()
	fmt.Println(st1)
}