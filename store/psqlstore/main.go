package psqlstore

import (
	"github.com/jackc/pgx/v5"
)

const (
	ErrSongNotFound = "song not found"
)

type PsqlStore struct {
	Conn *pgx.Conn
}

func New() (*PsqlStore, error) {
	conn, err := newConn()
	if err != nil {
		return nil, err
	}

	return &PsqlStore{
		Conn: conn,
	}, nil
}
