package main

import "fmt"

func main() {
c:=make(chan int )	
go recieveOnlyChannel(c)
for  v := range c {
	fmt.Println(v)
}
}
func recieveOnlyChannel(ch chan<-int){
	for i := 0; i < 100; i++ {
		ch<-i
	}
	close(ch)
}