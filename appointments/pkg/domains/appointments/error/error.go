package error

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
)

const (
	defaultResponse string = "Sorry, something went wrong"
	defaultCode     int    = http.StatusInternalServerError
)

var (
	// ErrNew represents an error when creating a new Quote
	ErrNew = errors.New("unable to create new quote")
	// ErrEmptyRepository repository cannot be nil
	ErrEmptyRepository = errors.New("empty repository")
	// ErrTypeAssertion arises while trying to perform interface{}.(T)
	ErrTypeAssertion  = errors.New("unable to execute type assertion")
	ErrNotFound       = errors.New("Appointment not found")
	ErrDatabase       = errors.New("An error happened in database")
	ErrMemoryDatabase = errors.New("An error happened in memory database")
	ErrInvalidPath    = errors.New("Cannot read path")
	ErrInvalidBody    = errors.New("Invalid body")
)

type errorResponse struct {
	response string
	code     int
}

type restError map[error]errorResponse

// RESTErrorBussines Errors you want to map to more meaning response for clients and set specific
// HTTP status code should be included here
var RESTErrorBussines = restError{
	ErrNew:            {"Sorry, we cannot create a new appointment", http.StatusInternalServerError},
	sql.ErrNoRows:     {"Record not found", http.StatusNotFound},
	ErrNotFound:       {"Appointment not found", http.StatusNotFound},
	ErrDatabase:       {"An error happened in database", http.StatusInternalServerError},
	ErrInvalidPath:    {"Cannot read path", http.StatusBadRequest},
	ErrInvalidBody:    {"Invalid body", http.StatusBadRequest},
	ErrMemoryDatabase: {"Memory Database error", http.StatusBadRequest},
}

func (re restError) ErrorProcess(err error) (string, int) {
	for rErr, resp := range re {
		if errors.Is(err, rErr) {
			return resp.response, resp.code
		}
	}

	return defaultResponse, defaultCode
}
