// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetLibraryDetailsLibraryErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetLibraryDetailsLibraryErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetLibraryDetailsLibraryErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetLibraryDetailsLibraryErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetLibraryDetailsLibraryResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetLibraryDetailsLibraryResponseBody struct {
	Errors []GetLibraryDetailsLibraryErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetLibraryDetailsLibraryResponseBody{}

func (e *GetLibraryDetailsLibraryResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetLibraryDetailsErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetLibraryDetailsErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetLibraryDetailsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetLibraryDetailsErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetLibraryDetailsResponseBody - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetLibraryDetailsResponseBody struct {
	Errors []GetLibraryDetailsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetLibraryDetailsResponseBody{}

func (e *GetLibraryDetailsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}