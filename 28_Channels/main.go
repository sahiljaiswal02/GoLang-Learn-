// Channels are the way by which multiple GORoutines communicate with each other.
package main

import "fmt"

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2)  // 2 refers to how may values can be stored in the channel or returning, which is also not necessary to give.
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {   // <-chan this is going outside the box so it is receiving the value
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen) // If before passing the value close() tag is used then it will return false and the value returned will be 0.
		fmt.Println(val) 
		// fmt.Println(<-myCh)   // 5 6
		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {  // chan<- this is going inside the box so it is sending the value
		myCh <- 5
		myCh <- 6	
		close(myCh)
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}