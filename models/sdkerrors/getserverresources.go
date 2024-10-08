// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetServerResourcesPlexErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetServerResourcesPlexErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetServerResourcesPlexErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetServerResourcesPlexErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetServerResourcesUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetServerResourcesUnauthorized struct {
	Errors []GetServerResourcesPlexErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetServerResourcesUnauthorized{}

func (e *GetServerResourcesUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetServerResourcesErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetServerResourcesErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetServerResourcesErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetServerResourcesErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetServerResourcesBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetServerResourcesBadRequest struct {
	Errors []GetServerResourcesErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetServerResourcesBadRequest{}

func (e *GetServerResourcesBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
