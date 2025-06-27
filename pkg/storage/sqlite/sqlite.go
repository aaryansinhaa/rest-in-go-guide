package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/aaryansinhaa/miyuki/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	DB *sql.DB
}

func NewSQLiteStorage(cfg *config.Config) (*SQLiteStorage, error) {
	storage, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Connecting to SQLite database at %s\n", cfg.StoragePath)
	_, err = storage.Exec(`CREATE TABLE IF NOT EXISTS blocks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL,
		address TEXT NOT NULL,
		email TEXT NOT NULL
	)`)
	if err != nil {
		return nil, err
	}
	return &SQLiteStorage{DB: storage}, nil

}

func (s *SQLiteStorage) CreateBlock(name string, age int, address string, email string) (int64, error) {
	//we used s.DB.Prepare instead of s.DB.Exec to prepare the statement
	// to avoid SQL injection attacks and to improve performance
	//we can use s.DB.Exec if we are not using the prepared statement
	//but it is not recommended
	//as it is less secure and less performant.
	result, err := s.DB.Prepare("INSERT INTO blocks (name, age, address, email) VALUES (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer result.Close()
	res, err := result.Exec(name, age, address, email)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Printf("Block created with ID: %d\n", id)
	return id, nil
}
