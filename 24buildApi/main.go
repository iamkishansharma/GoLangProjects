package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course - file
type Course struct {
	CourseId     string  `json:"courseid"`
	CourseName   string  `json:"coursename"`
	CoursePrice  float64 `json:"price"`
	CourseBanner string  `json:"banner"`

	Author *Author `json:"author"`
}

// Temp database of SLICE
var courses []Course

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Middleware or helper
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to Building API with GO")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{
		CourseId:     "1",
		CourseName:   "React Js Development",
		CourseBanner: "https://d29xdxvhssor07.cloudfront.net/assets/schools/2410/courses/53635/reactimage_fq7uz.png",
		CoursePrice:  299,
		Author: &Author{
			FullName: "Hitesh Choudahry",
			Website:  "kishan.com",
		},
	})
	courses = append(courses, Course{
		CourseId:     "2",
		CourseBanner: "https://d29xdxvhssor07.cloudfront.net/assets/schools/2410/courses/63685/angular_9q5mqo.png",
		CourseName:   "Angular Js Development",
		CoursePrice:  299,
		Author: &Author{
			FullName: "Hitesh Choudahry",
			Website:  "web.hiteshchoudhary.com",
		},
	})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getACourse).Methods("GET")
	r.HandleFunc("/course", creatingACourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateACourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteACourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// Controllers
// serving HomePage on Web
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Kishan Sharma...</h1>"))
}

// serving all courses as JSON on web
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllCourses triggered!")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// serve a course after searchin with given id
func getACourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetOneCourse triggered!")
	w.Header().Set("Content-Type", "application/json")

	// getting id
	params := mux.Vars(r)

	fmt.Printf("Type of params: %T\n", params)

	// finding same id in the DB (SLICE)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id.")
}

// serve a response after adding a course
func creatingACourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreatACourse triggered!")
	w.Header().Set("Content-Type", "application/json")

	// check for empty body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	// check for '{}' empty data
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No JSON data found.")
		return
	}
	// check if the course name is dublicate
	//TODO:: loop check is the same name is availabe and thrw error response

	// generate uid and convert to string
	// add to courses db (append to SLICE)

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode("Course successfully added")
}

// serve response after updating course
func updateACourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateACourse triggered!")
	w.Header().Set("Content-Type", "application/json")

	// from the is inside slice and remove that course
	// Again adding the new course with same id

	// grab id from web request
	params := mux.Vars(r)

	// finding same id in the DB ( or SLICE) to update
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// removing the course with same id
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			// adding new course
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Course successfully updated.")

			return
		}
	}

	//TODO:: send response if ID not found
}

//serve response after deleting a course with id
func deleteACourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteACourse triggered!")
	w.Header().Set("Content-Type", "application/json")

	// getting id
	params := mux.Vars(r)

	// finding same id in the DB (or SLICE) and deleting it
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// deleting course
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course successfully deleted.")
			return
		}
	}
	json.NewEncoder(w).Encode("Cannot delete course. ")
}
