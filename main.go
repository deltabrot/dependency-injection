package main

import (
	"log"

	"github.com/deltabrot/dependency-injection/handler"
	"github.com/deltabrot/dependency-injection/store/psqlstore"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	s, err := psqlstore.New()
	if err != nil {
		panic(err)
	}

	h := handler.New(s)
	log.Fatal(h.Run())
}
