package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error) {
	response := CreateNewResponse(400, "error", err.Error())
	log.Println(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
