package controllers

import (
	"aluraflix-api/models"
	"encoding/json"
	"net/http"
)

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	videos := models.ListarTodosOsVideos()
	json.NewEncoder(w).Encode(videos)
}
