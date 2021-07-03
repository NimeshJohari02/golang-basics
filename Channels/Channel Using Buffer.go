package main

import "fmt"

func main() {
	//! Here the 10 acts like a buffer and does cause a deadlock 
	ch:=make(chan int ,10);
	ch<-56;
	fmt.Println(<-ch)
	sendOnlyChannels:=make(chan <- int,3)
	sendOnlyChannels<-120
	sendOnlyChannels<-121
	sendOnlyChannels<-311
	// fmt.Println(<-sendOnlyChannels)As you see this lines Gives an Err 
	// recieveOnlyChannel:= make(<-chan int )
	//This is receive only 
	// fmt.Println(<-recieveOnlyChannel) This also Produces Error as empty receive only channel is created 
}