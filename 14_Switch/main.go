package main
import "fmt"
	   "math/rand"
	   "time"

func main() {
	fmt.Println("Switch in golang")

	rand.seed(time.Now().UnixNano())  //To generate random number

	diceNumber := rand.Intn(6) + 1  //To generate random number between 1-6
	fmt.Println("Value of dice is: ", diceNumber)  

	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
		fallthrough    //fallthrough will execute the case 3 as well as the case 4
	case 4:
		fmt.Println("Dice value is 4")
	case 5:
		fmt.Println("Dice value is 5")
	case 6:
		fmt.Println("Dice value is 6")
	default:
		fmt.Println("What was that?")	
	}
}