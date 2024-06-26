// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetResourcesStatisticsErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetResourcesStatisticsErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetResourcesStatisticsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetResourcesStatisticsErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetResourcesStatisticsResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetResourcesStatisticsResponseBody struct {
	Errors []GetResourcesStatisticsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetResourcesStatisticsResponseBody{}

func (e *GetResourcesStatisticsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
