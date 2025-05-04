Course Management System API
This is my RESTful API for managing students, courses, and enrollments. I wrote it in Go using gorilla/mux for routing. It handles CRUD operations for students and courses and manages student enrollments. Data is stored in memory.
How to Run the Project
Prerequisites

Go (v1.18 or higher)
curl or Postman for testing

Setup

Clone the repo:
git clone git@github.com:Turonsho/course-api.git
cd course-api


Install dependencies:
go mod tidy


Run the server:
go run main.go

The API runs at http://localhost:8080.


Notes

Data is in memory and resets on server restart.
I included sample data with a few students and courses.

API Overview
My API manages students, courses, and enrollments. It uses JSON for requests and responses. I used HTTP status codes: 200 for success, 404 for not found, 400 for bad requests.
Base URL
http://localhost:8080
Endpoints
Student Endpoints



Method
Endpoint
Description



GET
/students
Get all students


GET
/students/{id}
Get a student by ID


POST
/students
Create a new student


PUT
/students/{id}
Update a student


DELETE
/students/{id}
Delete a student


Course Endpoints



Method
Endpoint
Description



GET
/courses
Get all courses


GET
/courses/{id}
Get a course by ID


POST
/courses
Create a new course


PUT
/courses/{id}
Update a course


DELETE
/courses/{id}
Delete a course


Enrollment Endpoints



Method
Endpoint
Description



POST
/students/{id}/courses/{cid}
Enroll a student in a course


DELETE
/students/{id}/courses/{cid}
Unenroll a student from a course


GET
/students/{id}/courses
Get all courses for a student


GET
/courses/{id}/students
Get all students in a course


Sample Requests and Responses
Here are examples using curl. You can also use Postman.
1. Create a New Student
   Request:
   curl -X POST http://localhost:8080/students \
   -H "Content-Type: application/json" \
   -d '{"name": "Alice Smith", "email": "alice.smith@example.com"}'

Response:
{
"id": "3",
"name": "Alice Smith",
"email": "alice.smith@example.com"
}

Status: 201 Created
2. Get All Courses
   Request:
   curl http://localhost:8080/courses

Response:
[
{
"id": "1",
"title": "Web Development",
"description": "Learn to build web apps",
"instructor": "Dr. Brown"
},
{
"id": "2",
"title": "Machine Learning",
"description": "Intro to ML",
"instructor": "Prof. Lee"
}
]

Status: 200 OK
3. Enroll a Student in a Course
   Request:
   curl -X POST http://localhost:8080/students/3/courses/1

Response:
{
"message": "Student enrolled successfully"
}

Status: 201 Created
4. Get Courses for a Student
   Request:
   curl http://localhost:8080/students/3/courses

Response:
[
{
"id": "1",
"title": "Web Development",
"description": "Learn to build web apps",
"instructor": "Dr. Brown"
}
]

Status: 200 OK
5. Error: Student Not Found
   Request:
   curl http://localhost:8080/students/999

Response:
{
"error": "Student not found"
}

Status: 404 Not Found
Testing

Test endpoints with curl or Postman.
I handled errors for invalid inputs, non-existent students/courses, and duplicate enrollments.
Sample data loads on server start.

Notes

I used gorilla/mux for routing.
Data is in memory for simplicity. For production, I’d use a database.
I added validation for required fields and emails.

