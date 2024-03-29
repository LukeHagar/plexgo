// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

var GetTokenServerList = []string{
	"https://plex.tv/api/v2",
}

type GetTokenRequest struct {
	// The PinID to retrieve an access token for
	PinID string `pathParam:"style=simple,explode=false,name=pinID"`
	// The unique identifier for the client application
	// This is used to track the client application and its usage
	// (UUID, serial number, or other number unique per device)
	//
	XPlexClientIdentifier string `header:"style=simple,explode=false,name=X-Plex-Client-Identifier"`
}

func (o *GetTokenRequest) GetPinID() string {
	if o == nil {
		return ""
	}
	return o.PinID
}

func (o *GetTokenRequest) GetXPlexClientIdentifier() string {
	if o == nil {
		return ""
	}
	return o.XPlexClientIdentifier
}

type GetTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
