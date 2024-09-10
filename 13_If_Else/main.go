package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10"
	}
	fmt.Println(result)  // Watch out

	if 9%2 ==0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")  
	}
	// Output: Number is odd

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}
	// Output : Number is less than 10
}