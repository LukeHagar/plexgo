// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

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
	PlayQueueID *float64 `queryParam:"style=form,explode=true,name=playQueueID"`
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

func (o *AddPlaylistContentsRequest) GetPlayQueueID() *float64 {
	if o == nil {
		return nil
	}
	return o.PlayQueueID
}

type AddPlaylistContentsMetadata struct {
	RatingKey    *string `json:"ratingKey,omitempty"`
	Key          *string `json:"key,omitempty"`
	GUID         *string `json:"guid,omitempty"`
	Type         *string `json:"type,omitempty"`
	Title        *string `json:"title,omitempty"`
	Summary      *string `json:"summary,omitempty"`
	Smart        *bool   `json:"smart,omitempty"`
	PlaylistType *string `json:"playlistType,omitempty"`
	Composite    *string `json:"composite,omitempty"`
	Duration     *int    `json:"duration,omitempty"`
	LeafCount    *int    `json:"leafCount,omitempty"`
	AddedAt      *int    `json:"addedAt,omitempty"`
	UpdatedAt    *int    `json:"updatedAt,omitempty"`
}

func (o *AddPlaylistContentsMetadata) GetRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.RatingKey
}

func (o *AddPlaylistContentsMetadata) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *AddPlaylistContentsMetadata) GetGUID() *string {
	if o == nil {
		return nil
	}
	return o.GUID
}

func (o *AddPlaylistContentsMetadata) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AddPlaylistContentsMetadata) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *AddPlaylistContentsMetadata) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *AddPlaylistContentsMetadata) GetSmart() *bool {
	if o == nil {
		return nil
	}
	return o.Smart
}

func (o *AddPlaylistContentsMetadata) GetPlaylistType() *string {
	if o == nil {
		return nil
	}
	return o.PlaylistType
}

func (o *AddPlaylistContentsMetadata) GetComposite() *string {
	if o == nil {
		return nil
	}
	return o.Composite
}

func (o *AddPlaylistContentsMetadata) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *AddPlaylistContentsMetadata) GetLeafCount() *int {
	if o == nil {
		return nil
	}
	return o.LeafCount
}

func (o *AddPlaylistContentsMetadata) GetAddedAt() *int {
	if o == nil {
		return nil
	}
	return o.AddedAt
}

func (o *AddPlaylistContentsMetadata) GetUpdatedAt() *int {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type AddPlaylistContentsMediaContainer struct {
	Size               *int                          `json:"size,omitempty"`
	LeafCountAdded     *int                          `json:"leafCountAdded,omitempty"`
	LeafCountRequested *int                          `json:"leafCountRequested,omitempty"`
	Metadata           []AddPlaylistContentsMetadata `json:"Metadata,omitempty"`
}

func (o *AddPlaylistContentsMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *AddPlaylistContentsMediaContainer) GetLeafCountAdded() *int {
	if o == nil {
		return nil
	}
	return o.LeafCountAdded
}

func (o *AddPlaylistContentsMediaContainer) GetLeafCountRequested() *int {
	if o == nil {
		return nil
	}
	return o.LeafCountRequested
}

func (o *AddPlaylistContentsMediaContainer) GetMetadata() []AddPlaylistContentsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// AddPlaylistContentsResponseBody - Playlist Updated
type AddPlaylistContentsResponseBody struct {
	MediaContainer *AddPlaylistContentsMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *AddPlaylistContentsResponseBody) GetMediaContainer() *AddPlaylistContentsMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type AddPlaylistContentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Playlist Updated
	Object *AddPlaylistContentsResponseBody
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

func (o *AddPlaylistContentsResponse) GetObject() *AddPlaylistContentsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
