package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&payment=ghbj456ghb"

func main() {
	fmt.Println("Handling URL in golang")
	fmt.Println("URL is: ", myurl)

	//Parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)  //https
	fmt.Println(result.Host)  //lco.dev:3000
	fmt.Println(result.Path)  ///learn
	fmt.Println(result.Port())  //3000
	fmt.Println(result.RawQuery)  //coursename=reactjs&payment=ghbj456ghb

	qparams := result.Query()
	fmt.Printf("Type of qparams is: %T\n", qparams) //Type of qparams is: url.Values

	fmt.Println(qparams["coursename"])  //[reactjs]

	for _, val := range qparams{
		fmt.Println("Param is: ",val)  //Param is: [reactjs] Param is: [ghbj456ghb]
	}

	// Create URL
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=sahil",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)  //https://lco.dev/tutcss?user=sahil
}