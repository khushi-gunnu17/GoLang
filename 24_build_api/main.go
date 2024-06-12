package main

import (
	"encoding/json"
	"fmt"
	"log"
	// "math/rand"
	"net/http"
	// "strconv"
	// "time"
	"crypto/rand"
	"math/big"

	"github.com/gorilla/mux"
)

// Model for courses and author - should be in a separate file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"couresname"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course


// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}


func main() {

	fmt.Println("-------------------building-api-------------------")

	r := mux.NewRouter()

	// seeding of the data 
	courses = append(courses, Course{
		CourseId: "3", 
		CourseName : "ReactJS",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "Khushi Sharma",
			Website: "khushi.dev",
		},
	})

	courses = append(courses, Course{
		CourseId: "6", 
		CourseName : "MERN-Stack",
		CoursePrice: 199,
		Author: &Author{
			Fullname: "Khushi Sharma",
			Website: "go.dev",
		},
	})


	// routing
	r.HandleFunc("/", serveHome).Methods("Get")
	r.HandleFunc("/courses", getAllCourses).Methods("Get")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("Get")
	r.HandleFunc("/course/create", createOneCourse).Methods("POST")
	r.HandleFunc("/course/update/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/delete/{id}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses/deleteAll", deleteAllCourses).Methods("DELETE")


	// listening to a port
	log.Fatal(http.ListenAndServe(":5000", r))

}


// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to API</h1>"))

}


// controllers
func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)

}



func getOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get One Course")

	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)		// We can also use URL library here.

	// looping through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given ID.")
	// return	// redundant return statement

}



func createOneCourse(w http.ResponseWriter, r *http.Request)  {
	
	fmt.Println("Creating a course")

	w.Header().Set("Content-Type", "application/json")

	// checking if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	// what about : empty json given {}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON.")
		return
	}

	// checking if the title is duplicate
	// loop, title matches with course.coursename, then sending the JSON response
	for _, existingCourse := range courses {
		if existingCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course title already exists")
			return
		}
	}

	// generate a unique id and convert that into the string
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(100))
	course.CourseId = myRandomNum.String()

	// rand.Seed(time.Now().UnixNano())			// rand.Seed is deprecated, but we can also use this 
	// course.CourseId = strconv.Itoa(rand.Intn(100))		// Itoa = format Integer

	// appending course into courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	
}



func updateOneCourse(w http.ResponseWriter, r *http.Request)  {

	fmt.Println("Updating a course")

	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)
	
	// loop through the value, grab the Id and remove it from the Index, adding the value again with my ID
	for index, course := range courses {

		if course.CourseId == params["id"] {

			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		} 
	}

	// TODO : send a response when ID is not found
	json.NewEncoder(w).Encode("ID is not found")

}



func deleteOneCourse(w http.ResponseWriter, r *http.Request)  {
	
	fmt.Println("Deleting a Course")

	w.Header().Set("Content-Type", "application/json")

	// first grab id from the req
	params := mux.Vars(r)

	// looping thorugh the value and then grabbing the index and at last removing the index
	for index, course := range courses {

		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			// Sending a confirm message
			json.NewEncoder(w).Encode("Deleted the course")
			return
		}
	}

	// If the course ID is not found, sending an error message
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error" : "Course ID not found."})

}



// func deleteAllCourses()  {}	: complete it
func deleteAllCourses(w http.ResponseWriter, r *http.Request)  {

	fmt.Println("Deleting all the courses.")

	w.Header().Set("Content-Type", "application/json")

	// clearing the list of courses
	// The courses slice is reassigned to an empty slice, effectively clearing all the existing courses.
	courses = []Course{}

	// sending a confirmation message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted all courses"})
	
}



/*
	first ==> go build .
	then ==>  go run main.go
*/