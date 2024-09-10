package main
import "fmt"	

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)  //[Sunday Monday Tuesday Wednesday Thursday Friday Saturday]

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])  //Sunday Monday Tuesday Wednesday Thursday Friday Saturday
	}

	for i := range days {
		fmt.Println(days[i])  //Sunday Monday Tuesday Wednesday Thursday Friday Saturday
	}

	for index, day := range days{
		fmt.Printf("index is %v and value is %v\n", index, day) //index is 0 and value is Sunday  and so on ...
	}

	rogue := 1
	for rogue < 10 {

		if rogue == 2 {  // Output: 1 Hellooo
			goto lco
		}

		if rogue == 5{ // Output: 1 2 3 4
			break
		}  

		if rogue == 7{  // Output: 1 2 3 4 5 6 8 9
			rogue++
			continue
		}

		fmt.Println(rogue)  // 1 2 3 4 5 6 7 8 9
		rogue++
	}

	lco: 
		fmt.Println("Hellooo")
}