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
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React JS",
		"Price": 200,
		"website": "Udemy",
		"tags": ["web-dev","react"]
	}
	`)
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)  // &main.course{Name:"React JS",Price:200,Platform:"Udemy",Password:"12345",Tags:[]string{"web-dev","react"}}
	} else {
		fmt.Println("JSON was not valid")
	}

	//some case where you want key value pair
	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData) 
	fmt.Printf("%#v\n", myOnlineData) //map[coursename:React JS Price:200 website:Udemy tags:[web-dev react]] 


	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T\n", k, v, v) 
	}
	//Output: 
	//key is coursename and value is React JS and type is string key is Price and value is 200 and type is int64 key is website and value is Udemy and type is string key is tags and value is [web-dev react] and type is []string 
}
	