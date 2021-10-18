package router

import (
	"github.com/gorilla/mux"
	"github.com/iamkishansharma/GoLangProjects/tree/25ApiWithMongoDB/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovie).Methods("GET")
	// router.HandleFunc("/api/movie/{id}", controller.GetAMovie).Methods("GET")
	router.HandleFunc("/api/movie", controller.InsertAMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/movies/delete", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
