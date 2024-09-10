package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greet()  //Good Morning

	// Functions inside functions are not allowed in Go
	// func greet2() {
	// 	fmt.Println("Good Night")
	// }
	// greet2()

	result := add(2, 3)
	fmt.Println("Result is: ",result) // Result is:  5

	proRes, myMessage := proAdder(1, 2, 3, 4, 5)  //As we are returning two values so we have to handle both the values
	fmt.Println("Pro Result is: ",proRes) // Pro Result is:  15
	fmt.Println("My Message is: ",myMessage) // My Message is:  Hi i am proAdder function

}

func proAdder(values ...int) (int, string) {  //... is used for variable and take any number of arguments
	total = 0					// As we are returning two values so we have to wrap up them in braces

	for _, val := range values {
		total += val
	}
	return total, "Hi i am proAdder function"
}

func greet() {
	fmt.Println("Good Morning")
}

func add (a int, b int) int {
	return a + b	
}