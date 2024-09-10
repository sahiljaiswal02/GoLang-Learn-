package main

import "fmt"

func main(){
	fmt.Println("Methods in golang")

	sahil := User{"Sahil", "sahil@gmail.com", true, 16}
	fmt.Println(sahil) //{Sahil sahil@gmail.com true 16}
	fmt.Printf("Details of sahil are: %+v\n") //Details of sahil are: {Name: Sahil Email: sahil@gmail.com Status: true Age: 16}
	fmt.Printf("Name is : %v and Email is: %v\n", sahil.Name, sahil.Email) //Name is : Sahil and Email is: sahil@gmail.com

	sahil.GetStatus()  //Is user active:  true
	sahil.NewEmail()   //Email of this user is:  test@example.com
}

type User struct {  //All the names initialized with capital letter are public
	Name string
	Email string 
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail() {  // It doesn't change the original email, it just passes the copy of the email and changes it
	u.email = "test@example.com"
	fmt.Println("Email of this user is: ", u.Email)
} 