package controllers

import (
	"aluraflix-api/holders"
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(holders.Holder{Mensagem: "Bem Vindo"})
}
