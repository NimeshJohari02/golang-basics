package main

import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup
func f2(){
	for i := 0; i < 10; i++ {
		fmt.Println("f2 ",i)
	}
	wg.Done()
}
func f1(){
	for i := 0; i < 10; i++ {
		fmt.Println("f1", i)
	}
}
func genInfo(){
	fmt.Println("\nOS\t",runtime.GOOS)
	fmt.Println("\nARCH\t",runtime.GOARCH)
	fmt.Println("\nCPU\t",runtime.NumCPU())
	fmt.Println("\nRoutines\t",runtime.NumGoroutine())
}
func main() {
	genInfo();
	wg.Add(1)
	f1()
	go f2()
	genInfo()
	wg.Wait()
}