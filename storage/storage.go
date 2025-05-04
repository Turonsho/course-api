package storage

import "course-api/models"

var Students = make(map[int]models.Student)
var Courses = make(map[int]models.Course)
var Enrollments = []models.Enrollment{}

func InitData() {
	Students[1] = models.Student{ID: 1, Name: "Alice", Email: "alice@example.com"}
	Students[2] = models.Student{ID: 2, Name: "Bob", Email: "bob@example.com"}

	Courses[1] = models.Course{ID: 1, Name: "Math 101", Description: "Intro to Mathematics"}
	Courses[2] = models.Course{ID: 2, Name: "Physics 101", Description: "Intro to Physics"}
}
