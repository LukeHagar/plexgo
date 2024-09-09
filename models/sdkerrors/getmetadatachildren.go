// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetMetadataChildrenLibraryErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetMetadataChildrenLibraryErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetMetadataChildrenLibraryErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetMetadataChildrenLibraryErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetMetadataChildrenUnauthorized - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetMetadataChildrenUnauthorized struct {
	Errors []GetMetadataChildrenLibraryErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetMetadataChildrenUnauthorized{}

func (e *GetMetadataChildrenUnauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type GetMetadataChildrenErrors struct {
	Code    *int64  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetMetadataChildrenErrors) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetMetadataChildrenErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetMetadataChildrenErrors) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetMetadataChildrenBadRequest - Bad Request - A parameter was not specified, or was specified incorrectly.
type GetMetadataChildrenBadRequest struct {
	Errors []GetMetadataChildrenErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetMetadataChildrenBadRequest{}

func (e *GetMetadataChildrenBadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
