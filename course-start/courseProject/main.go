package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db

var courses []course

func (c *course) Isempty() bool {
	//return c.CourseId=="" &&  c.CourseName==""
	return c.CourseName == ""
}

func main() {

	r := chi.NewRouter()

	ser := &http.Server{
		Handler: r,
		Addr:    ":3000",
	}
	ser.ListenAndServe()
	r.get("/")
	courses = append(courses, course{CourseId: "1", CourseName: "React.js", CoursePrice: 299,
		Author: &Author{Fullname: "rakesh", Website: "help.com"}})
	r.get("/courses", GetAllCourses)
	r.get("/courses/{id}", getOnecourse)
	r.post("/createCourse", createonecourse)
	r.delete("/deleteCourse", deleteOneCourse)
	r.Put("/updateCourse", updateonecourse)

}

//controller file
//servehome route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> welcome to api of learnby self</h1>"))
}
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses ")
	w.Header().Set("COntent-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "appliation/json")
	params := chi.URLParam("id")
	for _, course := range courses {
		if course.CourseId == params {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found in this course")
	return
}

func createonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create onr course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send me some method")
	}
	var course course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.Isempty() {
		json.NewEncoder(w).Encode("no data inside the course")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}
func updateonecourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := chi.URLParam("id")
	for index, course := range courses {
		if course.CourseId == params {
			courses = append(courses[:index], courses[index+1:]...)
			var course course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.courseId = params
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam("id")
	for idx, course := range courses {
		if course.CourseId == params {
			courses = append(courses[:idx], courses[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("value remove ")

}
