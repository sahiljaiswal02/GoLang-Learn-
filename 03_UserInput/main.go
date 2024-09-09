package main

import 
	"fmt"
	"bufio"
	"os"

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)  // it will take input from the user
	fmt.Println("Enter the rating for our pizza:")

	// comma ok || err ok
	input, _ = reader.ReadString('\n')  //\n is used to terminate the input whe the user hits enter
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is: %T", input)
}