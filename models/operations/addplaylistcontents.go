// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AddPlaylistContentsRequest struct {
	// the ID of the playlist
	PlaylistID float64 `pathParam:"style=simple,explode=false,name=playlistID"`
	// the content URI for the playlist
	URI string `queryParam:"style=form,explode=true,name=uri"`
	// the play queue to add to a playlist
	PlayQueueID float64 `queryParam:"style=form,explode=true,name=playQueueID"`
}

func (o *AddPlaylistContentsRequest) GetPlaylistID() float64 {
	if o == nil {
		return 0.0
	}
	return o.PlaylistID
}

func (o *AddPlaylistContentsRequest) GetURI() string {
	if o == nil {
		return ""
	}
	return o.URI
}

func (o *AddPlaylistContentsRequest) GetPlayQueueID() float64 {
	if o == nil {
		return 0.0
	}
	return o.PlayQueueID
}

type AddPlaylistContentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AddPlaylistContentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddPlaylistContentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddPlaylistContentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
