package memstore

import (
	"errors"
	"sync"

	"github.com/deltabrot/dependency-injection/model"
	"github.com/google/uuid"
)

const (
	ErrSongNotFound = "song not found"
)

type MemStore struct {
	songs     map[string]model.Song
	songMutex sync.Mutex
}

func New() *MemStore {
	return &MemStore{
		songs: make(map[string]model.Song),
	}
}

func (s *MemStore) GetSong(id string) (model.Song, error) {
	s.songMutex.Lock()
	defer s.songMutex.Unlock()

	song, ok := s.songs[id]
	if !ok {
		return model.Song{}, errors.New(ErrSongNotFound)
	}
	return song, nil
}

func (s *MemStore) CreateSong(song *model.Song) error {
	s.songMutex.Lock()
	defer s.songMutex.Unlock()

	song.ID = uuid.New().String()
	s.songs[song.ID] = *song

	return nil
}

func (s *MemStore) SayHi() {
	println("Hi")
}
