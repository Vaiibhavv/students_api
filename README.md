# 📘 Students API

A simple RESTful API built using Golang that allows users to perform **CRUD** operations (Create, Read, Update, Delete) on a list of students. Data is persisted using an **SQLite** database.

---

## 🚀 Features

- Create a new student
- Get all students or a specific student by ID
- Update student details
- Delete a student
- SQLite integration for data persistence

---

## 🛠️ Tech Stack

- **Golang** – Core language
- **SQLite** – Lightweight relational database
- **net/http** – Built-in HTTP router
- **github.com/mattn/go-sqlite3** – SQLite driver for Golang

---

## 📁 Project Structure

students_api/
├── main.go
├── database/
│ └── db.go
├── models/
│ └── student.go
├── handlers/
│ └── student_handlers.go
├── go.mod
└── README.md

## Json
[
  {
    "id": 1,
    "name": "John Doe",
    "age": 22,
    "grade": "A"
  }
]


