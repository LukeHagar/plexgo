// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetSessionHistorySessionsErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetSessionHistorySessionsErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetSessionHistorySessionsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetSessionHistorySessionsErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetSessionHistoryUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetSessionHistoryUnauthorized struct {
	Errors []GetSessionHistorySessionsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetSessionHistoryUnauthorized{}

func (e *GetSessionHistoryUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetSessionHistoryErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetSessionHistoryErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetSessionHistoryErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetSessionHistoryErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetSessionHistoryBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetSessionHistoryBadRequest struct {
	Errors []GetSessionHistoryErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetSessionHistoryBadRequest{}

func (e *GetSessionHistoryBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
