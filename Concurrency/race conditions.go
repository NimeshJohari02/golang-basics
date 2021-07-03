package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter:=0
	gr:=40
	var wg sync.WaitGroup
	wg.Add(gr)
	fmt.Print("\tGo Routines\t",runtime.NumGoroutine())
	fmt.Println("\n CPUS\t",runtime.NumCPU())
	for i := 0; i < gr; i++ {
		go func(){
			v:=counter
			// time.Sleep(time.Second*3)
			runtime.Gosched()
			v++
			counter=v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Print("\n Go Routines\t",runtime.NumGoroutine())
	fmt.Println("\nCounter\t ",counter) 
}