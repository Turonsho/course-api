package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"course-api/models"
	"course-api/storage"
	"github.com/gorilla/mux"
)

// GET /students
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	var students []models.Student
	for _, student := range storage.Students {
		students = append(students, student)
	}
	json.NewEncoder(w).Encode(students)
}

// POST /students
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	json.NewDecoder(r.Body).Decode(&student)

	student.ID = len(storage.Students) + 1
	storage.Students[student.ID] = student

	json.NewEncoder(w).Encode(student)
}

// DELETE /students/{id}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	delete(storage.Students, id)
	w.WriteHeader(http.StatusNoContent)
}
