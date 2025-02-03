package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"restApi/internal/storage"

	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storagePath string) (*Storage, error) {

	if storagePath == "" {
		return nil, errors.New("storage path is empty")
	}

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) SaveURL(urlToSave string, alias string) error {

	statement, err := s.db.Prepare("INSERT INTO links(url, alias) VALUES(?, ?)")

	defer statement.Close()

	if err != nil {
		return err
	}

	_, err = statement.Exec(urlToSave, alias)

	if err != nil {
		if sqliteError, ok := err.(sqlite3.Error); ok && sqliteError.Code == sqlite3.ErrConstraint {
			return fmt.Errorf("%w", storage.ErrorURLExists)
		}
		return err
	}

	return nil

}

func (s *Storage) GetURL(alias string) (string, error) {

	urlResult := ""

	row, err := s.db.Query("SELECT url FROM links WHERE alias = ?", alias)

	defer row.Close()

	if err != nil {
		return urlResult, fmt.Errorf("%w", storage.ErrorUrlsNotFound)
	}

	for row.Next() {
		err = row.Scan(&urlResult)
		if err != nil {
			return urlResult, err
		}
	}

	return urlResult, nil

}

func (s *Storage) DeleteByAlias(alias string) error {

	statement, err := s.db.Prepare("DELETE FROM links WHERE alias = ?")

	defer statement.Close()

	if err != nil {
		return err
	}

	_, err = statement.Exec(alias)

	if err != nil {
		return err
	}

	return nil
}
