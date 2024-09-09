package main

import "fmt"

const Logintoken = "abcdef"  // Initializing the name with capital letter declare that it is a public variable

func main(){
	var name string = "Sahil"
	fmt.Println(name)  
	fmt.Printf("Variable is type of: %T \n", name)  

	var age int = 23
	fmt.Println(age)
	fmt.Printf("Variable is type of: %T \n", age)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is type of: %T \n", isLoggedIn)

	// default values and some alias
	var username string
	fmt.Println(username) // zero value : Go doesn't provide any garbage value to the empty variable
	fmt.Printf("Variable is type of: %T \n", username)
	username = "Sahil"
	fmt.Println(username)

	// implicit type
	var website = "bento.me/jaiswalsahil"  //lexar will automatically identify the type and give it
	fmt.Println(website)
	fmt.Printf("Variable is type of: %T \n", website)

	//no var style
	numberOfUsers := 4  //:= is called walrus syntax
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is type of: %T \n", numberOfUsers)

	//Public variable
	fmt.Println(Logintoken)		
	fmt.Printf("Variable is type of: %T \n", Logintoken)

}