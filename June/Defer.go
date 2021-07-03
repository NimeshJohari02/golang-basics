package main

import "fmt"
func second()  {
	fmt.Println("This is Seconds")
}
func first()  {
	fmt.Println("This is First")
}
func main() {
	
	// defer second()
	second()
	first()
	// A panic generally indicates a programmer error (e.g., attempting to access an index
// of an array that’s out of bounds, forgetting to initialize a map, etc.) or an exceptional
	defer func(){
		str := recover()
		fmt.Println(str)
	}()
	panic("This is panic Calll")
}