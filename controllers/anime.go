package controllers

import (
	"encoding/json"
	"net/http"
	"anime-backend/services"
	"github.com/gorilla/mux"
)

func GetAnime(w http.ResponseWriter, r *http.Request) {
	animeList := services.GetAnimeList() 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animeList)
}
func GetAnimeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	anime := services.GetAnimeByID(id)

	if anime == nil {
		http.Error(w, "Anime not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(anime)
}
