package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitList [4]string

	fruitList[0] = "Apple"	
	fruitList[1] = "Orange"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)  //[Apple Orange  Peach]  //Between apple and orange there is one space while between orange and peach there is two space which indicate that there is a gap
	fmt.Println("Fruit list is: ", len(fruitList)) //4

	var vegList = [3]string{"Potato", "beans", "mushroom"}	
	fmt.Println("Veg list is: ", vegList)  //[potato beans mushroom]
}