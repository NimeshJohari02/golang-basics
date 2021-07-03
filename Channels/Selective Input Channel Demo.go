package main

import "fmt"

func main() {
	c:=make(chan int)
	go foo(c)
	bar(c)
	fmt.Println("About To Exit")	
}
func foo(ch chan<- int){
	ch<-4100
}
func bar(ch <-chan int){
	fmt.Println(<-ch)
}