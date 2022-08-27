package helpers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CustomError struct{}

func (e CustomError) ApiError(w http.ResponseWriter, status int, message string) {
	error := make(map[string]string)

	error["message"] = message
	error["status"] = strconv.Itoa(status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}
