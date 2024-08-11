package main

import (
	"25mongodb/controller"
	"25mongodb/router"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("MongoDB CRUD operations in Go!")
	// Connect to MongoDB
	controller.InitDb()
	router := router.InitializeRoutes()

	fmt.Println("Server is running on port 4000!")
	http.ListenAndServe(":4000", router)
}
