package main

import (
	"database/sql"
	"errors"
	"log"

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
	stmt, err := s.db.Prepare("INSERT INTO youtube (title, description, thumbnailUrl, time) VALUES ($1,$2,$3,$4)")

	if err != nil {
		log.Fatal("Error while preparing db entry: ", err.Error())
		return err
	}

	for _, videoInfo := range videoInfos {

		_, err := stmt.Exec(videoInfo.Title, videoInfo.Description, videoInfo.ThumbnailUrl, videoInfo.DateTime)

		if err != nil {
			log.Fatal("Error while executing insert query into db: ", err)
			return err
		}
	}

	log.Println("Successfully inserted rows into db")
	return nil
}

func (s *PostgreStore) GetDetailsFromDbPaginated(pageNo int) ([]VideoInfo, error) {

	limit := 5
	offset := (pageNo - 1) * limit

	rows, err := s.db.Query("SELECT id, title, description, thumbnailUrl, time FROM youtube ORDER BY time DESC LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		log.Fatal("Error while retrieving paginated data from db: ", err)
		return nil, err
	}

	defer rows.Close()
	var videoInfos []VideoInfo

	for rows.Next() {
		var videoInfo VideoInfo

		err := rows.Scan(&videoInfo.Id, &videoInfo.Title, &videoInfo.Description, &videoInfo.ThumbnailUrl, &videoInfo.DateTime)
		if err != nil {
			log.Fatal("Error while scaning rows: ", err)
			return nil, err
		}

		log.Println("Video info: ", videoInfo)

		videoInfos = append(videoInfos, videoInfo)

	}

	return videoInfos, nil
}

func (s *PostgreStore) GetDetailsUsingTitleAndDescription(title string, description string) ([]VideoInfo, error) {
	if len(title) == 0 && len(description) == 0 {
		return nil, errors.New("Insert Valid title or description")
	}

	var rows *sql.Rows
	var err error

	if len(title) != 0 && len(description) != 0 {
		rows, err = s.db.Query("SELECT id, title, description, thumbnailUrl, time FROM youtube WHERE description LIKE '%' || $1 || '%' AND title LIKE '%' || $2 || '%'", description, title)
	} else if len(title) == 0 && len(description) != 0 {
		rows, err = s.db.Query("SELECT id, title, description, thumbnailUrl, time FROM youtube WHERE description LIKE '%' || $1 || '%'", description)
	} else {
		rows, err = s.db.Query("SELECT id, title, description, thumbnailUrl, time FROM youtube WHERE  title LIKE '%' || $1 || '%'", title)

	}

	if err != nil {
		log.Fatal("Error while retrieving data using description and title from db: ", err)
		return nil, err
	}

	defer rows.Close()

	var videoInfos []VideoInfo

	for rows.Next() {
		var videoInfo VideoInfo

		err := rows.Scan(&videoInfo.Id, &videoInfo.Title, &videoInfo.Description, &videoInfo.ThumbnailUrl, &videoInfo.DateTime)
		if err != nil {
			log.Fatal("Error while scaning rows for title description action: ", err)
			return nil, err
		}

		log.Println("Video info: ", videoInfo)

		videoInfos = append(videoInfos, videoInfo)

	}

	return videoInfos, nil
}
