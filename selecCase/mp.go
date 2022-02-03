package main

import (
	"fmt"
	"time"
)

// When a routine is communicating with several channels
// is very useful to use select in order to interact in an ordered way with all of the messages that are getting received

// Select\
// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
// https://go.dev/tour/concurrency/5

func doSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go doSomething(d1, c1, 1)
	go doSomething(d2, c2, 2)

	// fmt.Println(<-c1)
	// fmt.Println(<-c2)

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)

		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)

		}
	}
}
