package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("r1 Printing i = ", i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 23; i++ {
			fmt.Println("r2 Printing i = ", i)
		}
		wg.Done()
	}()
	wg.Wait()
}
