// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type StartTaskErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *StartTaskErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *StartTaskErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *StartTaskErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// StartTaskResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type StartTaskResponseBody struct {
	Errors []StartTaskErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &StartTaskResponseBody{}

func (e *StartTaskResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
