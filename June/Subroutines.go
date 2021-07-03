package main

//!NOTE ...  before the data type indicated variadic parameter
//! whereas ... after variable like arr... spreads the content of the variale
//! Note Variadic Parameters are the last paramater you should pass . So that you can pass 0 or more params if needed
import "fmt"
func main() {
	xs:=[]float64{12,31,21,32,33,4,5}
	total:=0.0;
	for  _,v := range xs{
	total+=v;
	}
	fmt.Println(add(xs...)/float64(len(xs)))
	m2,isEven:=MulBy2AndOddEven(46);
	fmt.Println(m2)
	fmt.Println(isEven)
}
//! Evrything in Golang is pass by value 
func add(args... float64) float64{
	total:=0.0
	for _,v:=range args{
		total+=v
	}
		return total
}
/*
You can return multiple things from a function in golang example below
 func <function_name>(<input parameter> <input parameter type>)(<return parameter sepearted by comma>){
	 function_body
 }
*/
func MulBy2AndOddEven(num int) (int , bool){
	n2:=num*2
	var isEven bool
	if num%2==0{
		isEven=true;
	}else
	{
		isEven=false
	}
	return n2,isEven
}