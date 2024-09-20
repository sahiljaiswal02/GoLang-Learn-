// Concurrency
// The ability of a program to perform multiple tasks at the same time but one task at a time.

// Parallelism 
// The ability of a program to perform multiple tasks at the same time but all at the same time.

// GoRoutines
// Managed by GO Run time
// Flexible Stack: 2kb
// Do not communicate by sharing memory, instead share memory by communicating.

// Thread
// Managed by OS
// Fixed Stack: 1mb

// Mutex: Protects access to a shared resource is a mutual exclusion lock.
// It helps in multiple GORoutines to access the same resource.

package main
import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup  // Pointer
var mut sync.Mutex  // Pointer

func main() {
	// greeter("Hello")  // 6 times Hello
	go greeter("Hello")  // Using go keyword to create go routine and it will run concurrently 
	greeter("World") // 6 times World

	websitelist := []string{
		"https://go.dev"
		"https://google.com"
		"https://fb.com"
		"https://github.com"
		"https://instagram.com"
	}

	for _, web := range websitelist {
		// getStatusCode(web)	
		// Output: 	
		// 200 Status Code for https://go.dev
		// 200 Status Code for https://google.com
		// 200 Status Code for https://fb.com
		// 200 Status Code for https://github.com
		// 200 Status Code for https://instagram.com

		go getStatusCode(web)
		wg.Add(1)  // It will signal that one go routine is added
	}

	wg.Wait()  // Wait till all go routines are done

	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)  // sleep for 3 milliseconds
		fmt.Println(s)
	}
}


func getStatusCode(endpoint string) {
	defer wg.Done()  // It will signal that one go routine is done

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS! in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d Status Code for %s\n", res.StatusCode, endpoint)
	}
}
