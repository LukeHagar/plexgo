// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type ApplyUpdatesUpdaterErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *ApplyUpdatesUpdaterErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ApplyUpdatesUpdaterErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ApplyUpdatesUpdaterErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// ApplyUpdatesUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type ApplyUpdatesUnauthorized struct {
	Errors []ApplyUpdatesUpdaterErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &ApplyUpdatesUnauthorized{}

func (e *ApplyUpdatesUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type ApplyUpdatesErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *ApplyUpdatesErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ApplyUpdatesErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ApplyUpdatesErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// ApplyUpdatesBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type ApplyUpdatesBadRequest struct {
	Errors []ApplyUpdatesErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &ApplyUpdatesBadRequest{}

func (e *ApplyUpdatesBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
