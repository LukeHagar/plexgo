// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/LukeHagar/plexgo/internal/utils"
	"net/http"
)

// IncludeDetails - Whether or not to include details for a section (types, filters, and sorts).
// Only exists for backwards compatibility, media providers other than the server libraries have it on always.
type IncludeDetails int64

const (
	IncludeDetailsZero IncludeDetails = 0
	IncludeDetailsOne  IncludeDetails = 1
)

func (e IncludeDetails) ToPointer() *IncludeDetails {
	return &e
}

func (e *IncludeDetails) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = IncludeDetails(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IncludeDetails: %v", v)
	}
}

type GetLibraryRequest struct {
	// the Id of the library to query
	SectionID float64 `pathParam:"style=simple,explode=false,name=sectionId"`
	// Whether or not to include details for a section (types, filters, and sorts).
	// Only exists for backwards compatibility, media providers other than the server libraries have it on always.
	//
	IncludeDetails *IncludeDetails `default:"0" queryParam:"style=form,explode=true,name=includeDetails"`
}

func (g GetLibraryRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetLibraryRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetLibraryRequest) GetSectionID() float64 {
	if o == nil {
		return 0.0
	}
	return o.SectionID
}

func (o *GetLibraryRequest) GetIncludeDetails() *IncludeDetails {
	if o == nil {
		return nil
	}
	return o.IncludeDetails
}

type GetLibraryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetLibraryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLibraryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLibraryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}