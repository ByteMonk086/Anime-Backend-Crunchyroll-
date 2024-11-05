package main

import (
	"fmt"
	"log"
	"net/http"
	"anime-backend/routes"
)

func main() {
	r := routes.SetupRoutes() 
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r)) 
}
