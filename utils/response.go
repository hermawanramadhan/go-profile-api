package utils

import (
	"customer-profile/entities"
	"encoding/json"
	"net/http"
)

func JsonSuccessResponse(w http.ResponseWriter, data interface{}, msg string) {

	result, _ := json.Marshal(entities.Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: msg,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func JsonErrorResponse(w http.ResponseWriter, statusCode int, msg string) {
	result, _ := json.Marshal(entities.Response{
		Code:    statusCode,
		Message: msg,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}
