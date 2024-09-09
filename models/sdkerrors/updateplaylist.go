// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type UpdatePlaylistPlaylistsErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *UpdatePlaylistPlaylistsErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *UpdatePlaylistPlaylistsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdatePlaylistPlaylistsErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// UpdatePlaylistUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type UpdatePlaylistUnauthorized struct {
	Errors []UpdatePlaylistPlaylistsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &UpdatePlaylistUnauthorized{}

func (e *UpdatePlaylistUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type UpdatePlaylistErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *UpdatePlaylistErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *UpdatePlaylistErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdatePlaylistErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// UpdatePlaylistBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type UpdatePlaylistBadRequest struct {
	Errors []UpdatePlaylistErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &UpdatePlaylistBadRequest{}

func (e *UpdatePlaylistBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
