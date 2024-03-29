package psqlstore

import (
	"context"

	"github.com/deltabrot/dependency-injection/model"
)

func (s *PsqlStore) GetSong(id string) (model.Song, error) {
	var song model.Song
	err := s.Conn.QueryRow(
		context.Background(),
		`
      SELECT
        id,
        title,
        artist
      FROM
        songs
      WHERE id = $1
      LIMIT 1
    `,
		id,
	).Scan(
		&song.ID,
		&song.Title,
		&song.Artist,
	)
	if err != nil {
		return model.Song{}, err
	}

	return song, nil
}

func (s *PsqlStore) CreateSong(song *model.Song) error {
	err := s.Conn.QueryRow(
		context.Background(),
		`
      INSERT INTO songs (
        title,
        artist
      ) VALUES (
        $1,
        $2
      ) RETURNING id
    `,
		song.Title,
		song.Artist,
	).Scan(&song.ID)
	if err != nil {
		return err
	}

	return nil
}
