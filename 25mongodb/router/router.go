package router

import (
	"25mongodb/controller"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movies", controller.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/movie", controller.InsertMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", controller.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{id}", controller.DeleteMovie).Methods("DELETE")

	return router
}
