package store

import "github.com/deltabrot/dependency-injection/model"

type Store interface {
	GetSong(string) (model.Song, error)
	CreateSong(*model.Song) error
}
