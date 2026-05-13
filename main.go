package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Message == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: call AI provider
	resp := ChatResponse{Reply: "echo: " + req.Message}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/chat", chatHandler)
	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
