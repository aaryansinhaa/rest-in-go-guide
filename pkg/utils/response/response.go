package response

import (
	"encoding/json"
	"net/http"

	"github.com/aaryansinhaa/miyuki/pkg/types"
)

func ResponseWriter(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}

func ErrorResponseWriter(w http.ResponseWriter, statusCode int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := types.ErrorResponse{
		Status:  statusCode,
		Message: message,
	}
	return json.NewEncoder(w).Encode(errorResponse)
}
