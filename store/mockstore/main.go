package mockstore

import (
	"github.com/deltabrot/dependency-injection/model"
)

type MockStore struct {
	response interface{}
	err      error
}

func New() *MockStore {
	return &MockStore{}
}

func (s *MockStore) SetResponse(model interface{}) {
	s.response = model
}

func (s *MockStore) SetError(err error) {
	s.err = err
}

func respond[T any](response interface{}) T {
	if res, ok := response.(T); ok {
		return res
	}
	var res T
	return res
}

func (s *MockStore) getError() error {
	err := s.err
	s.err = nil
	return err
}

func (s *MockStore) GetSong(id string) (model.Song, error) {
	return respond[model.Song](s.response), s.getError()
}

func (s *MockStore) CreateSong(song *model.Song) error {
	return s.getError()
}
