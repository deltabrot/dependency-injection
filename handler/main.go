package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/deltabrot/dependency-injection/model"
	"github.com/deltabrot/dependency-injection/store"
	"github.com/gorilla/mux"
)

type Handler struct {
	store  store.Store
	Router *mux.Router
}

func New(store store.Store) *Handler {
	h := &Handler{
		store: store,
	}

	router := mux.NewRouter()
	router.HandleFunc("/song/{id}", h.handleGetSong).Methods("GET")
	router.HandleFunc("/song", h.handleCreateSong).Methods("POST")

	h.Router = router

	return h
}

func (h *Handler) Run() error {
	http.Handle("/", h.Router)
	log.Println("Listening on :8080")
	return http.ListenAndServe(":8080", nil)
}

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
