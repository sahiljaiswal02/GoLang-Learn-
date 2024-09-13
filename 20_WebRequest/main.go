package main

import "fmt"
import "net/http"
import "io/ioutil"

const url = "https://lco.dev"
func main() {
	fmt.Println("Web request in golang")

	response, err := http.Get(url)

	if err != nil{
		panic(err)
	}
	fmt.Printf("Response is type of %T\n", response)  //Response is of type *http.Response
	defer response.Body.Close()  //caller's responsibility to close the connection
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}
	content := string(databytes)
	fmt.Println("Content is : ", content)  //You will the full body of the program 
}