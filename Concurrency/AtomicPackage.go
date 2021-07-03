package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
fmt.Println("Go Routines ", runtime.NumGoroutine())	
fmt.Println("CPUs ", runtime.NumCPU())	
var wg sync.WaitGroup
var counter int64
number_of_routine :=100;
wg.Add(number_of_routine)
for i := 0; i < number_of_routine; i++ {
	go func(){
		atomic.AddInt64(&counter,1);
		runtime.Gosched()	
		fmt.Println("Counter Value " , counter);
		wg.Done()
	}()
}
wg.Wait()
fmt.Println("Go Routines" , runtime.NumGoroutine())
fmt.Println("Counter Value  " ,counter)
}