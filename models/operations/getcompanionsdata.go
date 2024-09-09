// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

var GetCompanionsDataServerList = []string{
	"https://plex.tv/api/v2/",
}

type ResponseBody struct {
	Identifier string `json:"identifier"`
	BaseURL    string `json:"baseURL"`
	Title      string `json:"title"`
	LinkURL    string `json:"linkURL"`
	Provides   string `json:"provides"`
	// The plex authtoken used to identify with
	Token string `json:"token"`
}

func (o *ResponseBody) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *ResponseBody) GetBaseURL() string {
	if o == nil {
		return ""
	}
	return o.BaseURL
}

func (o *ResponseBody) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *ResponseBody) GetLinkURL() string {
	if o == nil {
		return ""
	}
	return o.LinkURL
}

func (o *ResponseBody) GetProvides() string {
	if o == nil {
		return ""
	}
	return o.Provides
}

func (o *ResponseBody) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type GetCompanionsDataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Companions Data
	ResponseBodies []ResponseBody
}

func (o *GetCompanionsDataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCompanionsDataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCompanionsDataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCompanionsDataResponse) GetResponseBodies() []ResponseBody {
	if o == nil {
		return nil
	}
	return o.ResponseBodies
}