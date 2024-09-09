// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type PostUsersSignInDataAuthenticationErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *PostUsersSignInDataAuthenticationErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *PostUsersSignInDataAuthenticationErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PostUsersSignInDataAuthenticationErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// PostUsersSignInDataUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type PostUsersSignInDataUnauthorized struct {
	Errors []PostUsersSignInDataAuthenticationErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &PostUsersSignInDataUnauthorized{}

func (e *PostUsersSignInDataUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type PostUsersSignInDataErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *PostUsersSignInDataErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *PostUsersSignInDataErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PostUsersSignInDataErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// PostUsersSignInDataBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type PostUsersSignInDataBadRequest struct {
	Errors []PostUsersSignInDataErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &PostUsersSignInDataBadRequest{}

func (e *PostUsersSignInDataBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
