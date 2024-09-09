package main

import "fmt"

func main() {
	fmt.Println("Pointers in golang")

	// var ptr *int  //Creating a pointer
	// fmt.Println("Value of pointer is: ", ptr)  //<nil>

	myNumber := 23
	var ptr = &myNumber  //Adding reference to the pointer
	fmt.Println("Value of actual pointer is: ", ptr)  //0xc00000a098
	fmt.Println("Value of pointer is: ", *ptr)  //23

	*ptr = *ptr + 2
	fmt.Println("Value of pointer is: ", myNumber)  //25


}

// NOTE:

// There are sometimes a copy of the variable is passed so to avoid it the pointers are used 
// so that it ensures that original value which is stored at the address is passed

// Just to create a pointer we use '*'
// but if we want to create a pointer with reference then we use '&'