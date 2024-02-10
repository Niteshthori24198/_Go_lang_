package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type Course struct {
	Title    string  `json:"title"`
	Price    int     `json:"price"`
	CourseId int     `json:"id"`
	Author   *Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var courses = []Course{{
	Title:    "Go",
	Price:    100,
	CourseId: 1,
	Author: &Author{
		Name:  "Nitesh",
		Email: "nitesh@gmail",
	},
},
	{
		Title:    "Java",
		Price:    10,
		CourseId: 2,
		Author: &Author{
			Name:  "Abhirup",
			Email: "abhirup@gmail",
		},
	},
}

func main() {

	// definning routes and controlling function

	http.HandleFunc("/", homeRoute)
	http.HandleFunc("/course", handleCourse)
	http.HandleFunc("/course/", handleCourse)

	// setting server

	log.Fatal(http.ListenAndServe(":6969", nil))

}

func homeRoute(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Welcome to home page💖")
}

func sayGoodByy(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("GoodBye👋")
}

// setting CRUD operations

func handleCourse(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		fmt.Println("Get course")
		if strings.HasPrefix(r.URL.Path, "/course/") {
			getOneCourseById(w, r)
		} else {
			getAllCourses(w, r)
		}

	case "POST":
		fmt.Println("Create course")
		createNewCourse(w, r)

	case "PUT":
		fmt.Println("Update course")
		updateCourseByID(w, r)

	case "DELETE":
		fmt.Println("Delete course")
		deleteCourseById(w, r)

	default:
		fmt.Println(" Opps Route Not found 🏏🧨")
		sayGoodByy(w, r)
	}

}

// fetching All courses informations

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

// fetch one specific course information

func getOneCourseById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// query := r.URL.Query()
	// fmt.Printf("params type %T\n ", query["id"][0])

	// taking it from params

	params := strings.Split(r.URL.Path, "/")
	courseID, _ := strconv.Atoi(strings.TrimSpace(params[2]))

	for _, course := range courses {

		if course.CourseId == courseID {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Course Not Found🤷‍♂️")
}

// create new course

func createNewCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var newCourse Course
	json.NewDecoder(r.Body).Decode(&newCourse)
	newCourse.CourseId = rand.Intn(100)

	courses = append(courses, newCourse)

	json.NewEncoder(w).Encode(newCourse)
}

// update existing course

func updateCourseByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := strings.Split(r.URL.Path, "/")

	courseID, _ := strconv.Atoi(strings.TrimSpace(params[2]))

	for index, course := range courses {

		if course.CourseId == courseID {

			var updatedCourse Course

			json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = courseID

			courses = append(courses[0:index], courses[index+1:]...)
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}

	json.NewEncoder(w).Encode("Course Not Found🤷‍♂️")
}

// delete existing course

func deleteCourseById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := strings.Split(r.URL.Path, "/")

	courseID, _ := strconv.Atoi(strings.TrimSpace(params[2]))

	for index, course := range courses {
		if course.CourseId == courseID {
			courses = append(courses[0:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted Successfully👍")
			return
		}
	}

	json.NewEncoder(w).Encode("Course Not Found🤷‍♂️")
}
