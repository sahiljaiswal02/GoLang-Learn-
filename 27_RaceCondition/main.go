package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition in golang")	

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("One")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("Two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("Three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
	
	// Output:
	// Race Condition in golang
	// Three
	// Two
	// One
	// [0 3 2 1]

}

// go run --race  // It will check weather the program is having any race condition or not.
// and if it gives you nothing then it means you are good to go, there's nothing wrong.