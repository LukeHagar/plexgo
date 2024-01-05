// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type StopTranscodeSessionErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *StopTranscodeSessionErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *StopTranscodeSessionErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *StopTranscodeSessionErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// StopTranscodeSessionResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type StopTranscodeSessionResponseBody struct {
	Errors []StopTranscodeSessionErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &StopTranscodeSessionResponseBody{}

func (e *StopTranscodeSessionResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}