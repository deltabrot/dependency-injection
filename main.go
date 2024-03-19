package main

import (
	"log"

	"github.com/deltabrot/dependency-injection/handler"
	"github.com/deltabrot/dependency-injection/model"
	"github.com/deltabrot/dependency-injection/store/mockstore"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	// var s store.Store
	// s, err := psqlstore.New()
	// if err != nil {
	//   panic(err)
	// }

	s := mockstore.New()

	s.SetResponse(model.Song{
		ID:     "1",
		Title:  "Hello",
		Artist: "Adele",
	})

	h := handler.New(s)
	log.Fatal(h.Run())
}
