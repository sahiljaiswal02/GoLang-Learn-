package main

import "fmt"

func main() {
	fmt.Println("Defer in golang")
	// The things that get passed in defer just get skipped and executed at the last of the program
	// And it uses LIFO structure for execution
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	//Output: Hello Two One World

	myDefer()

	//Output: Hello 4 3 2 1 0 Two One World
}


func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}