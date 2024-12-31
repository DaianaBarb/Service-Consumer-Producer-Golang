package utils

import (
	"encoding/json"
	"net/http"
	"project-golang/internal/utils/errors"
)

type ResponseDefault struct {
	Message string `json:"message"`
}

func ErrorResponse(w http.ResponseWriter, err error) {
	if errors.IsNotFound(err) {
		w.WriteHeader(http.StatusNotFound)
		encodeErr := json.NewEncoder(w).Encode(ResponseDefault{Message: err.Error()})
		if encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	if errors.IsBadRequest(err) {
		w.WriteHeader(http.StatusBadRequest)
		encodeErr := json.NewEncoder(w).Encode(ResponseDefault{Message: err.Error()})
		if encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	if errors.IsAlreadyExists(err) {
		w.WriteHeader(http.StatusConflict)
		encodeErr := json.NewEncoder(w).Encode(ResponseDefault{Message: err.Error()})
		if encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	if errors.IsConflict(err) {
		w.WriteHeader(http.StatusConflict)
		encodeErr := json.NewEncoder(w).Encode(ResponseDefault{Message: err.Error()})
		if encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	if errors.IsUnprocessableEntity(err) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		encodeErr := json.NewEncoder(w).Encode(ResponseDefault{Message: err.Error()})
		if encodeErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(ResponseDefault{Message: "internal server error: " + err.Error()})
}

func SuccessResponse(w http.ResponseWriter, statusCode int, res interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
