// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type StartUniversalTranscodeErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *StartUniversalTranscodeErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *StartUniversalTranscodeErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *StartUniversalTranscodeErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// StartUniversalTranscodeResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type StartUniversalTranscodeResponseBody struct {
	Errors []StartUniversalTranscodeErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &StartUniversalTranscodeResponseBody{}

func (e *StartUniversalTranscodeResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
