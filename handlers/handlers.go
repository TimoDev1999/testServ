package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"testServ/brokers"
	"testServ/database"
	"testServ/models"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	message.Processed = false
	if err := database.SaveMessage(&message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := brokers.ProduceMessage(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}
func ProcessedMessagesStats(w http.ResponseWriter, r *http.Request) {
	count, err := database.GetProcessedMessagesCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"processed_messages_count": count})
}
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/messages", CreateMessage).Methods("POST")
	r.HandleFunc("/messages/stats", ProcessedMessagesStats).Methods("GET")
	return r
}
