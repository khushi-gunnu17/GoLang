package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)


var signals = []string{"test"}

// variable which is going to be a wait group
// usually these are pointers, but here in this case we are not doing that 
var wg sync.WaitGroup

var mut sync.Mutex	// should be a pointer 

func main() {

	fmt.Println("------------------Go-Routines--------------------")
	fmt.Println("")

	// go greeter("Hello")
	// greeter("world")

	websiteList := []string {
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		wg.Add(1)
		go getStatusCode(web)
	}

	// responsible for not allowing the main method to finish, until and unless the go routines are completed.
	wg.Wait()

	fmt.Println("")
	fmt.Println(signals)

}


// func greeter(s string)  {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(4 * time.Millisecond)
// 		fmt.Println(s)
// 	}	
// }



func getStatusCode(endpoint string)  {

	defer wg.Done()

	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("something wrong in the endpoint.")
	} else {

		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s.\n", result.StatusCode, endpoint)
	}
	
}


// Mutex = Mutual Exclusion Lock = It provides you a lock over the memory.

// its purpose is that it locks the memory till that one go routine is working and till the time is writing anything inside that, it doesn't allow anybody else to use the memory. 






/*
Key Methods of sync.WaitGroup := 
- Add(int): Increments the WaitGroup counter by the specified number. Typically used to specify the number of goroutines that need to be waited for.
- Done(): Decrements the WaitGroup counter by one. Each goroutine should call this method when it finishes its work.
- Wait(): Blocks until the WaitGroup counter becomes zero. This is typically called by the main goroutine to wait for all other goroutines to complete.
*/