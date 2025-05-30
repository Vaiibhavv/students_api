package storage

import "githum.com/Vaiibhavv/students-api/students_api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	UpdateStudentDetails(id int64, name string, email string, age int) error
	DeleteStudentById(id int64) error
}
