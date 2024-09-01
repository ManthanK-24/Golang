package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myCh := make(chan int, 2)
    wg := &sync.WaitGroup{}

	// Deadlock Condition
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)

	// Receive Only Channel
	go func(ch <-chan int, wg *sync.WaitGroup){
		val, isChannelOpen := <- myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh,wg)

	// Send Only Channel
	go func(ch chan<- int, wg *sync.WaitGroup){
		// myCh <- 0
        close(myCh)
		//myCh <- 6
		wg.Done()
	}(myCh,wg)
    
	wg.Wait()
}