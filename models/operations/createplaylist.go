// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Type - type of playlist to create
type Type string

const (
	TypeAudio Type = "audio"
	TypeVideo Type = "video"
	TypePhoto Type = "photo"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "audio":
		fallthrough
	case "video":
		fallthrough
	case "photo":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

// Smart - whether the playlist is smart or not
type Smart int64

const (
	SmartZero Smart = 0
	SmartOne  Smart = 1
)

func (e Smart) ToPointer() *Smart {
	return &e
}

func (e *Smart) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = Smart(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Smart: %v", v)
	}
}

type CreatePlaylistRequest struct {
	// name of the playlist
	Title string `queryParam:"style=form,explode=true,name=title"`
	// type of playlist to create
	Type Type `queryParam:"style=form,explode=true,name=type"`
	// whether the playlist is smart or not
	Smart Smart `queryParam:"style=form,explode=true,name=smart"`
	// the content URI for the playlist
	URI *string `queryParam:"style=form,explode=true,name=uri"`
	// the play queue to copy to a playlist
	PlayQueueID *float64 `queryParam:"style=form,explode=true,name=playQueueID"`
}

func (o *CreatePlaylistRequest) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *CreatePlaylistRequest) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}

func (o *CreatePlaylistRequest) GetSmart() Smart {
	if o == nil {
		return Smart(0)
	}
	return o.Smart
}

func (o *CreatePlaylistRequest) GetURI() *string {
	if o == nil {
		return nil
	}
	return o.URI
}

func (o *CreatePlaylistRequest) GetPlayQueueID() *float64 {
	if o == nil {
		return nil
	}
	return o.PlayQueueID
}

type CreatePlaylistResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePlaylistResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePlaylistResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePlaylistResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}