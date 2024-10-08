// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Tonight - Indicate that you want the update to run during the next Butler execution. Omitting this or setting it to false indicates that the update should install
type Tonight int64

const (
	TonightZero Tonight = 0
	TonightOne  Tonight = 1
)

func (e Tonight) ToPointer() *Tonight {
	return &e
}
func (e *Tonight) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = Tonight(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Tonight: %v", v)
	}
}

// Skip - Indicate that the latest version should be marked as skipped. The [Release] entry for this version will have the `state` set to `skipped`.
type Skip int64

const (
	SkipZero Skip = 0
	SkipOne  Skip = 1
)

func (e Skip) ToPointer() *Skip {
	return &e
}
func (e *Skip) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = Skip(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Skip: %v", v)
	}
}

type ApplyUpdatesRequest struct {
	// Indicate that you want the update to run during the next Butler execution. Omitting this or setting it to false indicates that the update should install
	Tonight *Tonight `queryParam:"style=form,explode=true,name=tonight"`
	// Indicate that the latest version should be marked as skipped. The [Release] entry for this version will have the `state` set to `skipped`.
	Skip *Skip `queryParam:"style=form,explode=true,name=skip"`
}

func (o *ApplyUpdatesRequest) GetTonight() *Tonight {
	if o == nil {
		return nil
	}
	return o.Tonight
}

func (o *ApplyUpdatesRequest) GetSkip() *Skip {
	if o == nil {
		return nil
	}
	return o.Skip
}

type ApplyUpdatesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ApplyUpdatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ApplyUpdatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ApplyUpdatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
