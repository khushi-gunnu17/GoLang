package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("---------------channels-----------------")

	ch := make(chan int, 2)		// 2 => buffered channel (int)
	wg := &sync.WaitGroup{}

	// value is going inside the box (channel)
	// ch <- 5

	// value is going outside the box (channel)
	// fmt.Println( <-ch )


	wg.Add(2)
	// channel responsible for receiving the value
	go func(ch <-chan int, wg *sync.WaitGroup) {		// Receive only

		val, isChannelOpen :=  <-ch

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println( <-ch )

		wg.Done()
		
	}(ch, wg)

	// channel responsible for putting the values into the channel.
	go func(ch chan<- int, wg *sync.WaitGroup) {		// Send only

		// ch <- 0
		ch <- 10
		ch <- 20

		close(ch)

		wg.Done()

	}(ch, wg)


	wg.Wait()

}


// In Go, channels are a powerful feature used for communication between goroutines. They allow you to send and receive values between different parts of your program running concurrently, ensuring safe and synchronized communication.

// channela are bi-directional butcan be made uni-directional