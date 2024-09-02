package utils

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(message string, result interface{}, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"Message":   message,
		"Data":      result,
		"Status":    http.StatusOK,
		"IsSuccess": true,
	})
}

func BadRequestResponse(error error, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"Message":   error.Error(),
		"Data":      nil,
		"Status":    http.StatusBadRequest,
		"IsSuccess": false,
	})
}

func ErrorResponse(error error, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"Message":   error.Error(),
		"Data":      nil,
		"Status":    http.StatusInternalServerError,
		"IsSuccess": false,
	})
}

func UnauthorizedResponse(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"Message":   "Unauthorized Request!",
		"Data":      nil,
		"Status":    http.StatusUnauthorized,
		"IsSuccess": false,
	})
}

func SchemaErrorResponse(msg string, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(res).Encode(map[string]interface{}{
		"Message":   msg,
		"Data":      nil,
		"Status":    http.StatusUnprocessableEntity,
		"IsSuccess": false,
	})
}

// NotFoundResponse sends a not found JSON response
func NotFoundResponse(message string, w http.ResponseWriter) {
	response := map[string]interface{}{
		"message": message,
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}
