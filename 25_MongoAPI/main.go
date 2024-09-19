package main

import (
	"fmt"
)

// import gorilla "github.com/gorilla/mux"
// import mongodb "go.mongodb.org/mongo-driver/mongo"

func main() {
	fmt.Println("MongoDB in golang")
	r := router.Router()	
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000")
}