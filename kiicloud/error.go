package kiicloud

import (
	"fmt"
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
	return fmt.Sprintf("ErrorResponse:%d:%s:%s", er.StatusCode, er.Code,
		er.Message)
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
