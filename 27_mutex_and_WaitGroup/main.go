package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("------------------Race Condition------------------")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}


	wg.Add(4)

	// IIFE = Immediately Invoked function execution
	go func(wg *sync.WaitGroup, mut *sync.RWMutex){

		fmt.Println("First routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()

	}(wg, mut)


	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex){

		fmt.Println("Second Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
		
	}(wg, mut)


	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex){

		fmt.Println("Third Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
		
	}(wg, mut)



	go func(wg *sync.WaitGroup, mut *sync.RWMutex){

		fmt.Println("Fourth Routine")

		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		
		wg.Done()
		
	}(wg, mut)


	wg.Wait()
	

	fmt.Println(score)

}



// go run --race .    ==> gives all the info regarding the race conditions
// not running here, coz I have 32 bit MinGW installed

// whenever the race flag gives you nothing, that says nothing is going wrong.