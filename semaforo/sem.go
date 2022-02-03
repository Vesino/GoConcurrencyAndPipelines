package main

import (
	"fmt"
	"sync"
	"time"
)

// Buffered channels as semaforos can be use in order to allow us a limited number of executions
// and avoid the execution for a number of undefined executions

func main() {
	c := make(chan int, 2)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait()
}

func doSomething(id int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Started ID %d\n", id)
	time.Sleep(4 * time.Second)
	fmt.Printf("End ID %d\n", id)
	<-c
}
