// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type CheckForUpdatesUpdaterErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *CheckForUpdatesUpdaterErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *CheckForUpdatesUpdaterErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CheckForUpdatesUpdaterErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// CheckForUpdatesUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type CheckForUpdatesUnauthorized struct {
	Errors []CheckForUpdatesUpdaterErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &CheckForUpdatesUnauthorized{}

func (e *CheckForUpdatesUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type CheckForUpdatesErrors struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

func (o *CheckForUpdatesErrors) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *CheckForUpdatesErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CheckForUpdatesErrors) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// CheckForUpdatesBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type CheckForUpdatesBadRequest struct {
	Errors []CheckForUpdatesErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &CheckForUpdatesBadRequest{}

func (e *CheckForUpdatesBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
