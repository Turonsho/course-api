package routes

import (
	"course-api/handlers"
	"course-api/storage"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	storage.InitData() // Инициализация памяти

	r := mux.NewRouter()

	// 🔹 Студенты
	r.HandleFunc("/students", handlers.GetAllStudents).Methods("GET")
	r.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")

	// 🔹 Курсы
	r.HandleFunc("/courses", handlers.GetAllCourses).Methods("GET")
	r.HandleFunc("/courses", handlers.CreateCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", handlers.DeleteCourse).Methods("DELETE")

	// 🔹 Зачисления
	r.HandleFunc("/enrollments", handlers.EnrollStudent).Methods("POST")
	r.HandleFunc("/students/{id}/courses", handlers.GetStudentCourses).Methods("GET")
	r.HandleFunc("/students/{id}/courses/{cid}", handlers.EnrollStudentToCourse).Methods("POST")
	r.HandleFunc("/students/{id}/courses/{cid}", handlers.RemoveStudentFromCourse).Methods("DELETE")
	r.HandleFunc("/courses/{id}/students", handlers.GetCourseStudents).Methods("GET")

	return r
}
