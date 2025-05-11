package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"githum.com/Vaiibhavv/students-api/students_api/internal/config"
	"githum.com/Vaiibhavv/students-api/students_api/internal/types"
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

func (s *SQLite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id,name,email,age from students where id=? limit 1")
	if err != nil {
		return types.Student{}, nil
	}

	defer stmt.Close()
	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no value found for id %s", fmt.Sprint(id))
		}
		return types.Student{}, err
	}

	return student, nil

}

// update students details sql query

func (s *SQLite) UpdateStudentDetails(id int64, name string, email string, age int) error {

	stmt, err := s.Db.Prepare("update students set name=?,email=?,age=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, email, age, id)
	return err
}
