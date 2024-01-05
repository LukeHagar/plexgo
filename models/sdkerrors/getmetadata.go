// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetMetadataErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetMetadataErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetMetadataErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetMetadataErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetMetadataResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetMetadataResponseBody struct {
	Errors []GetMetadataErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetMetadataResponseBody{}

func (e *GetMetadataResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}