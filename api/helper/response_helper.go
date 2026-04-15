package helper

import (
	"encoding/json"
	"net/http"
)

func ToJson(w http.ResponseWriter, statusCode int, message string, data any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"statusCode": statusCode,
		"message":    message,
		"data":       data,
	})
}
