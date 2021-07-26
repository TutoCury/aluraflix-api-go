package routes

import (
	"aluraflix-api/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/videos", controllers.GetAllVideos)
}
