package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 10; i++ {
		go doSomething(i, &wg)
	}
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}
