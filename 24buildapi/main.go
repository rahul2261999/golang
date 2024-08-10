package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Models

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake Db
var courses []Course

// middlewares, helpers

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Hello Golang REST API")
	courses = append(
		courses,
		Course{CourseId: "1", CourseName: "Golang Fund", CoursePrice: 200, Author: &Author{Fullname: "John Doe", Website: "johndoe.com"}},
		Course{CourseId: "2", CourseName: "Python Programming", CoursePrice: 150, Author: &Author{Fullname: "Jane Smith", Website: "janesmith.com"}},
	)

	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course", createCourse).Methods("POST")
	router.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	router.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", router))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, Golang Web Development!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course by ID")

	params := mux.Vars(r)

	for _, item := range courses {
		if item.CourseId == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&Course{})
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Request body is empty"})

		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode(map[string]string{"message": "Course Id and Name are required"})

		return
	}

	course.CourseId = strconv.FormatInt(time.Now().UnixNano(), 10)

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Update course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)

			fmt.Println(courses)

			json.NewEncoder(w).Encode("Course updated successfully")

			return
		}
	}

	json.NewEncoder(w).Encode("Course not found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode("Course deleted successfully")
}
