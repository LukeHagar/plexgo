// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type StartAllTasksButlerErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *StartAllTasksButlerErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *StartAllTasksButlerErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *StartAllTasksButlerErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// StartAllTasksUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type StartAllTasksUnauthorized struct {
	Errors []StartAllTasksButlerErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &StartAllTasksUnauthorized{}

func (e *StartAllTasksUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type StartAllTasksErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *StartAllTasksErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *StartAllTasksErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *StartAllTasksErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// StartAllTasksBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type StartAllTasksBadRequest struct {
	Errors []StartAllTasksErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &StartAllTasksBadRequest{}

func (e *StartAllTasksBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
