package main

import (
	"fmt"
	"time"
	"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Math in Golang")

	var numberOne int = 2
	var numberTwo float64 = 4.5

	// fmt.Println("The sum is : ", numberOne + int(numberTwo))  // This is not a valid syntax.
	// And by converting float to int will round off the decimal part and we will not get the accurate value.

	//Random Number
	rand.seed(time.Now().UnixNano())  // rand.seed generate random number but generation is done through some algo so something that changes every time is time so we pass the value time inside.
	fmt.Println(rand.Intn(5)) // This will generate the random number within 5.

	// Random from Crypto : This is more accurate because it is generated from cryptography.
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}