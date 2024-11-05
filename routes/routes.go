package routes

import (
	"anime-backend/controllers"
	"github.com/gorilla/mux"
)
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/anime", controllers.GetAnime).Methods("GET")
	r.HandleFunc("/api/anime/{id}", controllers.GetAnimeByID).Methods("GET")
	return r
}
