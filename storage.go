package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	AddInDb(videoInfo []VideoInfo) error
	GetDetailsFromDbPaginated(pageNo int) ([]VideoInfo, error)
	GetDetailsUsingTitleAndDescription(title string, description string) ([]VideoInfo, error)
}

type PostgreStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgreStore, error) {

	connStr := "user=postgres dbname=postgres password=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreStore{
		db: db,
	}, nil

}

func (s *PostgreStore) Init() error {
	return s.createTable()

}

func (s *PostgreStore) createTable() error {
	query := `create table if not exists youtube (
		id serial primary key,
		title varchar(255),
		description varchar(255),
		thumbnailUrl varchar(255),
		time timestamp

	)`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgreStore) AddInDb(videoInfos []VideoInfo) error {

	return nil
}

func (s *PostgreStore) GetDetailsFromDbPaginated(pageNo int) ([]VideoInfo, error) {
	return nil, nil
}

func (s *PostgreStore) GetDetailsUsingTitleAndDescription(title string, description string) ([]VideoInfo, error) {
	return nil, nil
}
