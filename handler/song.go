package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/deltabrot/dependency-injection/model"
	"github.com/gorilla/mux"
)

func (h *Handler) handleSong() {
	h.Router.HandleFunc("/song/{id}", h.handleGetSong).Methods("GET")
	h.Router.HandleFunc("/song", h.handleCreateSong).Methods("POST")
}

// handleGetSong is a http.HandlerFunc that retrieves a song from the store
// and writes it to the response.
func (h *Handler) handleGetSong(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	song, err := h.store.GetSong(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(song)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleCreateSong is a http.HandlerFunc that creates a song in the store
// and writes it to the response.
func (h *Handler) handleCreateSong(w http.ResponseWriter, r *http.Request) {
	var song model.Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.store.CreateSong(&song)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(song)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
