// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetStatisticsErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetStatisticsErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetStatisticsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetStatisticsErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetStatisticsResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetStatisticsResponseBody struct {
	Errors []GetStatisticsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetStatisticsResponseBody{}

func (e *GetStatisticsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
