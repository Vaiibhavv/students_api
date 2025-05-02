package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"githum.com/Vaiibhavv/students-api/students_api/internal/config"
)

type SQLite struct {
	Db *sql.DB
}

// here we are using the config as pointer. we dont want to copy , referencing and dereferencing
func New(cfg *config.Config) (*SQLite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Students (

	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &SQLite{
		Db: db,
	}, nil

}

func (s *SQLite) CreateStudent(name string, email string, age int) (int64, error) {
	// ? = prevent from sql injection
	stmt, err := s.Db.Prepare("INSERT INTO Students (name,email,age) values(?,?,?)")
	if err != nil {
		return 0, nil
	}
	// need to close after opening
	defer stmt.Close()

	//need to execute the provided sql query, remeber the sequence of fields(name, email, age)
	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, nil
	}

	// we will take the Id as last autoincrement id
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	//if everythings good, we will return the id and error as nil
	return lastId, nil

}
