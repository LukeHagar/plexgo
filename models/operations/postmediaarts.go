// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostMediaArtsRequest struct {
	// the id of the library item to return the posters of.
	RatingKey int64 `pathParam:"style=simple,explode=false,name=ratingKey"`
	// The URL of the image, if uploading a remote image
	URL *string `queryParam:"style=form,explode=true,name=url"`
	// The contents of the image, if uploading a local file
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	RequestBody *any `request:"mediaType=image/*"`
}

func (o *PostMediaArtsRequest) GetRatingKey() int64 {
	if o == nil {
		return 0
	}
	return o.RatingKey
}

func (o *PostMediaArtsRequest) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *PostMediaArtsRequest) GetRequestBody() *any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PostMediaArtsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostMediaArtsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostMediaArtsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostMediaArtsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
