package request

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// PayloadValidation is an interface used to validate the payload of a request.
type PayloadValidation interface {
	Validate(request *http.Request) error
}

// DecodeAndValidateJSONPayload decodes the body of the request into an object
// implementing the PayloadValidation interface.
// It also triggers a validation of object content.
func DecodeAndValidateJSONPayload(request *http.Request, v PayloadValidation) error {
	if err := json.NewDecoder(request.Body).Decode(v); err != nil {
		return err
	}
	return v.Validate(request)
}

// GetPayload decodes the body of the request into an object implementing the PayloadValidation interface.
func GetPayload[T PayloadValidation](r *http.Request) (T, error) {
	var payload T
	err := DecodeAndValidateJSONPayload(r, payload)
	if err != nil {
		return payload, errors.WithMessage(err, "Invalid request payload")
	}
	return payload, nil
}
