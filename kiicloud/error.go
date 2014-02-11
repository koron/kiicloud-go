package kiicloud

import (
	"net/http"
)

// ErrorResponse represent Kii Cloud's error JSON.
type ErrorResponse struct {
	StatusCode int
	Code       string `json:"errorCode"`
	Message    string `json:"message"`
}

// Error returns string format of ErrorResponse.
func (er *ErrorResponse) Error() string {
	return "KiiCloud Error:" + er.Code + ":" + string(er.StatusCode) + ":" + er.Message
}

func ToError(resp *http.Response) error {
	var er ErrorResponse
	er.StatusCode = resp.StatusCode
	err := parseJson(resp.Body, &er)
	if err != nil {
		return err
	}
	return &er
}
