// Package response provides convenience functions to write into a http.ResponseWriter.
package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	httperror "github.com/portainer/libhttp/error"
)

// JSON encodes data to rw in JSON format. Returns a pointer to a
// HandlerError if encoding fails.
func JSON(rw http.ResponseWriter, data interface{}) *httperror.HandlerError {
	rw.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		return httperror.InternalServerError("Unable to write JSON response", err)
	}

	return nil
}

// JSON encodes data to rw in YAML format. Returns a pointer to a
// HandlerError if encoding fails.
func YAML(rw http.ResponseWriter, data interface{}) *httperror.HandlerError {
	rw.Header().Set("Content-Type", "text/yaml")

	strData, ok := data.(string)
	if !ok {
		return httperror.InternalServerError("Unable to write YAML response", errors.New("failed to convert input to string"))
	}

	fmt.Fprint(rw, strData)

	return nil
}

// Empty merely sets the response code to NoContent (204).
func Empty(rw http.ResponseWriter) *httperror.HandlerError {
	rw.WriteHeader(http.StatusNoContent)
	return nil
}
