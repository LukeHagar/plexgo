// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetTranscodeSessionsSessionsErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetTranscodeSessionsSessionsErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetTranscodeSessionsSessionsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetTranscodeSessionsSessionsErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetTranscodeSessionsUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetTranscodeSessionsUnauthorized struct {
	Errors []GetTranscodeSessionsSessionsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetTranscodeSessionsUnauthorized{}

func (e *GetTranscodeSessionsUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetTranscodeSessionsErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetTranscodeSessionsErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetTranscodeSessionsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetTranscodeSessionsErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetTranscodeSessionsBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetTranscodeSessionsBadRequest struct {
	Errors []GetTranscodeSessionsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetTranscodeSessionsBadRequest{}

func (e *GetTranscodeSessionsBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
