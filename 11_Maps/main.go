package main

import "fmt"

func main() {

	languages = make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all the languages: ",languages) //List of all the languages:  map[JS:Javascript PY:Python RB:Ruby]
	fmt.Println("JS stands for: ",languages["JS"]) //JS stands for:  Javascript

	delete(languages, "RB")
	fmt.Println("List of all the languages: ",languages) //List of all the languages:  map[JS:Javascript PY:Python]

	//Iterate over map using for Loop
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)  //For key JS, value is Javascript //For key PY, value is Python
	}
}