package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("MOD in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

//Output: If everything works fien and go build is executed then it will be hosted on the localserver:4000

// go.mod //this is fiel is created by go mod init mod
// go mod tidy //to tidy up the mod
// go get -u github.com/gorilla/mux  //Download this modeule to get access to the package
// go mod verify  // to check the modules are correct

func greeter() {
	fmt.Println("Hey There !!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my World</h1>"))
}


