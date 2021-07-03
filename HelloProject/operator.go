package main

import "fmt"

//Global Declarations
var glo = 12

// To make global declarations and NOT INITIALISE THEM USE THIS
//var z int
// this assign zero value for a particular type of
func main() {
	fmt.Println(":= is short declaration operator that declares as well as stores the value in the variable")
	//this declare and initialize at the same time
	x := 20
	fmt.Println(x)
	/*The Difference between using var and := is the var can be used for global Declarations whereas the := can only be used inside the function */
	var y = 7
	fmt.Println(y)
	checkglo()
}

// The below function demonstrates that glo is a global variable.
func checkglo() {
	fmt.Println(glo)
}
