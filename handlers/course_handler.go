package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"course-api/models"
	"course-api/storage"
	"github.com/gorilla/mux"
)

// GET /courses
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	var courses []models.Course
	for _, course := range storage.Courses {
		courses = append(courses, course)
	}
	json.NewEncoder(w).Encode(courses)
}

// POST /courses
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	json.NewDecoder(r.Body).Decode(&course)

	course.ID = len(storage.Courses) + 1
	storage.Courses[course.ID] = course

	json.NewEncoder(w).Encode(course)
}

// DELETE /courses/{id}
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	delete(storage.Courses, id)
	w.WriteHeader(http.StatusNoContent)
}
