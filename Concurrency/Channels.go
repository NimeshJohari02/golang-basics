package main

import "fmt"
func f1 (x int ) int {
	return x*x
}

func main() {
	ch:=make(chan int)
	go func ()  {
		ch <- f1(25)
	}()
fmt.Println(<-ch)	
}