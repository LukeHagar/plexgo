// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UpdatePlaylistRequest struct {
	// the ID of the playlist
	PlaylistID float64 `pathParam:"style=simple,explode=false,name=playlistID"`
	// name of the playlist
	Title *string `queryParam:"style=form,explode=true,name=title"`
	// summary description of the playlist
	Summary *string `queryParam:"style=form,explode=true,name=summary"`
}

func (o *UpdatePlaylistRequest) GetPlaylistID() float64 {
	if o == nil {
		return 0.0
	}
	return o.PlaylistID
}

func (o *UpdatePlaylistRequest) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *UpdatePlaylistRequest) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

type UpdatePlaylistResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdatePlaylistResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePlaylistResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePlaylistResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
