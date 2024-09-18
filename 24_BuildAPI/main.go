package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux")
// importing go get -u github.com/gorilla/mux

// Model for course - file

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice float `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`	
	Website string `json:"website"`
}

// Fake DB
var courses []Course	

// middleware helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

//MAIN Function
func main() {
	fmt.Println("Building API in golang")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{
		CourseId: "1",
		CourseName: "React JS",
		CoursePrice: 200,
		Author: &Author{
			Fullname: "Sahil",
			Website: "bento.me/jaiswalsahil",
		}
	})
	courses = append(courses, Course{
		CourseId: "2",
		CourseName: "Angular JS",
		CoursePrice: 300,
		Author: &Author{
			Fullname: "Harry",
			Website: "go.dev",
		}
	})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}


// controllers - file

//serve Home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	// loop through courses, finding matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with the Id")
	return 
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// data like this - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No Data Inside JSON")
		return
	}

	// TODO: Check only if title is duplicate
	// loop, title, matches with course.coursename, JSON

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course")
	w.Header().Set("Content-Type", "application/json")

	// first grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)	
			course.CourseId = params["id"]
			course = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TODO: send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// loop, id, remove, (index, index+1)	
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or decline response
			break
		}
	}
}