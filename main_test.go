package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deltabrot/dependency-injection/handler"
	"github.com/deltabrot/dependency-injection/model"
	"github.com/deltabrot/dependency-injection/store/mockstore"
)

func TestGetSongHandler(t *testing.T) {
	s := mockstore.New()
	s.SetResponse(model.Song{
		ID:     "1",
		Title:  "Hello",
		Artist: "Adele",
	})

	h := handler.New(s)

	req, err := http.NewRequest("GET", "/song/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"id":"1","title":"Hello","artist":"Adele"}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
