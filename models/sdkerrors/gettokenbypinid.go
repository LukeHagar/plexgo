// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetTokenByPinIDPlexErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

func (o *GetTokenByPinIDPlexErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetTokenByPinIDPlexErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// GetTokenByPinIDResponseBody - Not Found or Expired
type GetTokenByPinIDResponseBody struct {
	Errors []GetTokenByPinIDPlexErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetTokenByPinIDResponseBody{}

func (e *GetTokenByPinIDResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetTokenByPinIDErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *GetTokenByPinIDErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetTokenByPinIDErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetTokenByPinIDErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetTokenByPinIDBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetTokenByPinIDBadRequest struct {
	Errors []GetTokenByPinIDErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetTokenByPinIDBadRequest{}

func (e *GetTokenByPinIDBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
