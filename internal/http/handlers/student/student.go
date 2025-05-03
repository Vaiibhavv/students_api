package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"githum.com/Vaiibhavv/students-api/students_api/internal/response"
	"githum.com/Vaiibhavv/students-api/students_api/internal/storage"
	"githum.com/Vaiibhavv/students-api/students_api/internal/types"
)

// later we are passing the storage for jsonn record to be stored in database
func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating a student")

		/*when we send the json data in body, to get that data first we need to decode that data in
		go lang*/
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// if error is not above (EOF) then print actual error is coming
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Now validate the request body ( json body)
		if err := validator.New().Struct(student); err != nil {
			valerrs := err.(validator.ValidationErrors) // type casting for handling slice vali func
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(valerrs))
			return
		}

		// need to call the createstudent method to create a student in database

		lastid, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created", slog.String("userid", fmt.Sprint(lastid)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		// print the response after success
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastid})

		//w.Write([]byte("welcome to student api"))
	}
}

// to get the students data from storage by id

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id") // "id"  should be similar to the endpoint
		slog.Info("Get students by id", slog.String("Id", id))

		intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, err := storage.GetStudentById(intid)
		if err != nil {
			slog.Error("error getting user %s", slog.String("id", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}
