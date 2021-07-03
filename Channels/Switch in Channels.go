package main

import "fmt"

func main() {
	even:=make(chan int)	
	odd:=make(chan int)	
	quit:=make(chan int)
	//sending 
	go func(even, odd ,quit chan<- int) {
		for i := 0; i < 100; i++ {
			if i%2==0 {
				even<-i
			}else{
				odd<-i
			}
		}
		close(even)
		close(odd)
		quit<-0;
	}(even,odd,quit)
	//receive 
	func(even , odd , quit <-chan int) {
		for{
			 select {
			 case v:=<-even:
				fmt.Println("From The Eve Channel :",v)
			 
			 case v:=<-odd:
				fmt.Println("From The Odd Channel :",v)
			case v:=<-quit:
				fmt.Println("From Quit Channel :",v)
				return;
		}
	}
	}(even,odd,quit)
	fmt.Println("About To Exit ")
}