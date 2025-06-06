// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostMediaPosterRequest struct {
	// the id of the library item to return the posters of.
	RatingKey int64 `pathParam:"style=simple,explode=false,name=ratingKey"`
	// The URL of the image, if uploading a remote image
	URL *string `queryParam:"style=form,explode=true,name=url"`
	// The contents of the image, if uploading a local file
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	RequestBody *any `request:"mediaType=image/*"`
}

func (o *PostMediaPosterRequest) GetRatingKey() int64 {
	if o == nil {
		return 0
	}
	return o.RatingKey
}

func (o *PostMediaPosterRequest) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *PostMediaPosterRequest) GetRequestBody() *any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PostMediaPosterResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostMediaPosterResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostMediaPosterResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostMediaPosterResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
