package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type course struct {
	Name  string `json:"coursename"`  //The name between backticks is the name which will be displayed
	Price   int
	Platform  string `json:"website"`
	Password string `json:"-"`  // the password will not be displayed
	Tags []string `json:"tags,omitempty"`  // omitempty will not display if the field is empty
}

func main() {
	fmt.Println("Json in golang")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React JS", 200, "Udemy", "12345", []string{"web-dev", "react"}},
		{"MERN JS", 300, "Udemy", "12345", []string{"full-stack", "Mern"}},
		{"Angular JS", 400, "Udemy", "12345", nil},
	}

	//package the data
	// finalJson := json.Marshal(lcoCourses)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", finalJson) 
	// Output: {"Name":"React JS","Price":200,"Platform":"Udemy","Password":"12345","Tags":["web-dev","react"]} {"Name":"MERN JS","Price":300,"Platform":"Udemy","Password":"12345","Tags":["full-stack","Mern"]} {"Name":"Angular JS","Price":400,"Platform":"Udemy","Password":"12345","Tags":null}
	

	//To orgainize the data more clean we use MarshalIndent
	finalJson := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson) 
    
}