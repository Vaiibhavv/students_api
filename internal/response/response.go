package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	// if we want to provide status and error in lowercase  , used struct serialize
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk   = "Okay"
	StatuError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	// need to set the header for input data
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)

	// here we need to encode input  data (reverse)
	return json.NewEncoder(w).Encode(data) // this will also handled the error

}

func GeneralError(err error) Response {
	return Response{
		Status: StatuError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {

	var errMsg []string

	for _, err := range errs {
		switch err.ActualTag() {

		case "required":
			errMsg = append(errMsg, fmt.Sprintf("Field %s is required field", err.Field()))
		default:
			errMsg = append(errMsg, fmt.Sprintf("field %s is invalid", err.Field()))
		}

	}
	return Response{
		Status: StatuError,
		Error:  strings.Join(errMsg, ", "),
	}
}
