package handler

import (
	"log"
	"net/http"

	"github.com/deltabrot/dependency-injection/store"
	"github.com/gorilla/mux"
)

type Handler struct {
	store  store.Store
	Router *mux.Router
}

// New returns a new Handler with the given store.
func New(store store.Store) *Handler {
	// initialise handler
	h := &Handler{
		store: store,
	}

	// initialise router
	h.Router = mux.NewRouter()

	// middleware
	h.Router.Use(ContentTypeJson)

	// routes
	h.handleSong()

	return h
}

// Run starts the http server and listens on port 8080.
func (h *Handler) Run() error {
	http.Handle("/", h.Router)
	log.Println("Listening on :8080")
	return http.ListenAndServe(":8080", nil)
}

// ContentTypeJson is a middleware function that when used causes all
// responses to have a content type of application/json.
func ContentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
