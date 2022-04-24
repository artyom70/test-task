package internal

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	// ErrBadRequest represents bad request error
	ErrBadRequest = errors.New("bad request")
	// ErrInternalError represents internal error
	ErrInternalError = errors.New("internal server error")
	// ErrNotFound represents not found error
	ErrNotFound = errors.New("not found")
	// ErrTimeout denotes a timeout error
	ErrTimeout = errors.New("timeout")
)

// EncodeError encodes error errors
func EncodeError(err error, w http.ResponseWriter) {
	switch {
	case errors.Is(err, ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, ErrBadRequest):
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
