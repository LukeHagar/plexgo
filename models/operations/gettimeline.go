// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

// State - The state of the media item
type State string

const (
	StatePlaying State = "playing"
	StatePaused  State = "paused"
	StateStopped State = "stopped"
)

func (e State) ToPointer() *State {
	return &e
}

type GetTimelineRequest struct {
	// The rating key of the media item
	RatingKey float64 `queryParam:"style=form,explode=true,name=ratingKey"`
	// The key of the media item to get the timeline for
	Key string `queryParam:"style=form,explode=true,name=key"`
	// The state of the media item
	State State `queryParam:"style=form,explode=true,name=state"`
	// Whether the media item has MDE
	HasMDE float64 `queryParam:"style=form,explode=true,name=hasMDE"`
	// The time of the media item
	Time float64 `queryParam:"style=form,explode=true,name=time"`
	// The duration of the media item
	Duration float64 `queryParam:"style=form,explode=true,name=duration"`
	// The context of the media item
	Context string `queryParam:"style=form,explode=true,name=context"`
	// The play queue item ID of the media item
	PlayQueueItemID float64 `queryParam:"style=form,explode=true,name=playQueueItemID"`
	// The playback time of the media item
	PlayBackTime float64 `queryParam:"style=form,explode=true,name=playBackTime"`
	// The row of the media item
	Row float64 `queryParam:"style=form,explode=true,name=row"`
}

func (o *GetTimelineRequest) GetRatingKey() float64 {
	if o == nil {
		return 0.0
	}
	return o.RatingKey
}

func (o *GetTimelineRequest) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *GetTimelineRequest) GetState() State {
	if o == nil {
		return State("")
	}
	return o.State
}

func (o *GetTimelineRequest) GetHasMDE() float64 {
	if o == nil {
		return 0.0
	}
	return o.HasMDE
}

func (o *GetTimelineRequest) GetTime() float64 {
	if o == nil {
		return 0.0
	}
	return o.Time
}

func (o *GetTimelineRequest) GetDuration() float64 {
	if o == nil {
		return 0.0
	}
	return o.Duration
}

func (o *GetTimelineRequest) GetContext() string {
	if o == nil {
		return ""
	}
	return o.Context
}

func (o *GetTimelineRequest) GetPlayQueueItemID() float64 {
	if o == nil {
		return 0.0
	}
	return o.PlayQueueItemID
}

func (o *GetTimelineRequest) GetPlayBackTime() float64 {
	if o == nil {
		return 0.0
	}
	return o.PlayBackTime
}

func (o *GetTimelineRequest) GetRow() float64 {
	if o == nil {
		return 0.0
	}
	return o.Row
}

type GetTimelineResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetTimelineResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTimelineResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTimelineResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
