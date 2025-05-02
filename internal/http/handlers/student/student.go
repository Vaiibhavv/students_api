package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"githum.com/Vaiibhavv/students-api/students_api/internal/response"
	"githum.com/Vaiibhavv/students-api/students_api/internal/types"
)

func New() http.HandlerFunc {
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
		}

		// print the response after success
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "okay"})
		//w.Write([]byte("welcome to student api"))
	}
}
