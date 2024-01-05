// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetButlerTasksErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetButlerTasksErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetButlerTasksErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetButlerTasksErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetButlerTasksResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetButlerTasksResponseBody struct {
	Errors []GetButlerTasksErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetButlerTasksResponseBody{}

func (e *GetButlerTasksResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}