// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetFileHashLibraryErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetFileHashLibraryErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetFileHashLibraryErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFileHashLibraryErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetFileHashUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetFileHashUnauthorized struct {
	Errors []GetFileHashLibraryErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetFileHashUnauthorized{}

func (e *GetFileHashUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetFileHashErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetFileHashErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetFileHashErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFileHashErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetFileHashBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetFileHashBadRequest struct {
	Errors []GetFileHashErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetFileHashBadRequest{}

func (e *GetFileHashBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
