package main

import (
	"course-api/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.SetupRoutes()
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
