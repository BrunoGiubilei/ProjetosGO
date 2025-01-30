package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brunogiubilei/go-rest-api/greetings"
)

func GetAPI(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func GetGreeting(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing 'name' query parameter", http.StatusBadRequest)
		return
	}

	greeting := greetings.GetGreeting(name)
	response := map[string]string{"message": greeting}
	json.NewEncoder(w).Encode(response)
}
