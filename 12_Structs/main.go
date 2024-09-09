package main

import "fmt"

func main(){
	fmt.Println("Structs in golang")

	//no inheritance and super or parent in golang	
	sahil := User{"Sahil", "sahil@gmail.com", true, 16}
	fmt.Println(sahil) //{Sahil sahil@gmail.com true 16}
	fmt.Printf("Details of sahil are: %+v\n") //Details of sahil are: {Name: Sahil Email: sahil@gmail.com Status: true Age: 16}
	fmt.Printf("Name is : %v and Email is: %v", sahil.Name, sahil.Email) //Name is : Sahil and Email is: sahil@gmail.com
}

type User struct {  //All the names initialized with capital letter are public
	Name string
	Email string 
	Status bool
	Age int
}