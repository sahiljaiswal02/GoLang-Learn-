package main

import "fmt"
	   "os"
	   "io"
	   "io/ioutil"

func main() {
	fmt.Println("Files in golang")

	content := "This needs to go in a file - sahiljaiswal"
	file, err := os.Create("./myfile.txt")
	checkNillErr(err)

	length, err := file.WriteString(file, content)
	checkNillErr(err)
	fmt.Println("Length is: ", length)
	defer file.close()  // defer is used here so that if there is any changes furthur then it gets executed and file gets closed at last

	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	checkNillErr(err)
	fmt.Println(string(databyte))

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}