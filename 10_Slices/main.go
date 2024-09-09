package main
import "fmt"

func main() {
	fmt.Println("Welcome to slices")

	var fruitArr = []string {"Apple", "Mango", "Melon"}
	fmt.Printf("The type of fruitList is %T\n", fruitList)  //[]string

	fruitList = append(fruitList, "Banana", "Grapes")
	fmt.Println(fruitList)  //[Apple Mango Melon Banana Grapes]

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList) //[Mango Melon Banana Grapes]

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList) //[Mango Melon]

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 455
	highScores[3] = 567

	fmt.Println(highScores) //[234 345 455 567]

	highScores = append(highScores, 555, 676, 778)
	fmt.Println(highScores) //[234 345 455 567 555 676 778]

	sort.Ints(highScores)
	fmt.Println(highScores) //[234 345 455 555 567 676 778]
	fmt.Println(sort.IntsAreSorted(highScores)) //true

	//how to remove values from slices based on index

	var courses = []string {"react", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses) //[react javascript swift python ruby]
	
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) //[react javascript python ruby]
}