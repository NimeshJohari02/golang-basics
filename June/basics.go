package main

import (
	"fmt"
	"math"
	"time"
)
func main(){
	fmt.Println(math.Sin(50000));
	fmt.Println("1+1",1+1);
	var a = 12;
	fmt.Print(a);
// Loops in GO 
for j:=1;j<=12;j++{
	fmt.Println("Value of j is",j);
}
//if else in go
if 7%2 == 0 {
	fmt.Println("7 is even ")
}
// No ternary Operator in golang 
i:=2;
switch i {
case 1:
	fmt.Println("one")
case 2: 
}

fmt.Println("two");

 t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    };
// ! Swirtch can also be used for various types 

whatami:= func (i interface{}) {
	switch t:=i.(type) {
	case bool:
		fmt.Print("Boolean Type")
	case int:
		fmt.Print("Interger Type")
	default:
	fmt.Printf("Don't know Type %T \n ",t);
}
}
whatami(23);
whatami(true);
x:=42
y:=45
fmt.Println("Currently x=",x,"y = ",y)
x,y=y,x
fmt.Println("Currently x=",x,"y = ",y)
}