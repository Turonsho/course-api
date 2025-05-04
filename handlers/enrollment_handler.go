package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"course-api/models"
	"course-api/storage"
	"github.com/gorilla/mux"
)

// ✅ POST /enrollments (старый способ)
func EnrollStudent(w http.ResponseWriter, r *http.Request) {
	var enrollment models.Enrollment
	json.NewDecoder(r.Body).Decode(&enrollment)

	storage.Enrollments = append(storage.Enrollments, enrollment)
	json.NewEncoder(w).Encode(enrollment)
}

// ✅ GET /students/{id}/courses
func GetStudentCourses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	studentID, _ := strconv.Atoi(params["id"])
	var studentCourses []models.Course

	for _, e := range storage.Enrollments {
		if e.StudentID == studentID {
			if course, ok := storage.Courses[e.CourseID]; ok {
				studentCourses = append(studentCourses, course)
			}
		}
	}

	json.NewEncoder(w).Encode(studentCourses)
}

// ✅ POST /students/{id}/courses/{cid}
func EnrollStudentToCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID, _ := strconv.Atoi(vars["id"])
	courseID, _ := strconv.Atoi(vars["cid"])

	enrollment := models.Enrollment{
		StudentID: studentID,
		CourseID:  courseID,
	}

	// Проверка на дубликаты
	for _, e := range storage.Enrollments {
		if e.StudentID == studentID && e.CourseID == courseID {
			http.Error(w, "Student already enrolled", http.StatusConflict)
			return
		}
	}

	storage.Enrollments = append(storage.Enrollments, enrollment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(enrollment)
}

// ✅ DELETE /students/{id}/courses/{cid}
func RemoveStudentFromCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID, _ := strconv.Atoi(vars["id"])
	courseID, _ := strconv.Atoi(vars["cid"])

	for i, e := range storage.Enrollments {
		if e.StudentID == studentID && e.CourseID == courseID {
			storage.Enrollments = append(storage.Enrollments[:i], storage.Enrollments[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Enrollment not found", http.StatusNotFound)
}

// ✅ GET /courses/{id}/students
func GetCourseStudents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseID, _ := strconv.Atoi(vars["id"])
	var courseStudents []models.Student

	for _, e := range storage.Enrollments {
		if e.CourseID == courseID {
			if student, ok := storage.Students[e.StudentID]; ok {
				courseStudents = append(courseStudents, student)
			}
		}
	}

	json.NewEncoder(w).Encode(courseStudents)
}
