// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/LukeHagar/plexgo/internal/utils"
	"github.com/LukeHagar/plexgo/types"
	"net/http"
)

// GetTopWatchedContentQueryParamType - The type of media to retrieve or filter by.
// 1 = movie
// 2 = show
// 3 = season
// 4 = episode
// E.g. A movie library will not return anything with type 3 as there are no seasons for movie libraries
type GetTopWatchedContentQueryParamType int64

const (
	GetTopWatchedContentQueryParamTypeMovie   GetTopWatchedContentQueryParamType = 1
	GetTopWatchedContentQueryParamTypeTvShow  GetTopWatchedContentQueryParamType = 2
	GetTopWatchedContentQueryParamTypeSeason  GetTopWatchedContentQueryParamType = 3
	GetTopWatchedContentQueryParamTypeEpisode GetTopWatchedContentQueryParamType = 4
	GetTopWatchedContentQueryParamTypeAudio   GetTopWatchedContentQueryParamType = 8
	GetTopWatchedContentQueryParamTypeAlbum   GetTopWatchedContentQueryParamType = 9
	GetTopWatchedContentQueryParamTypeTrack   GetTopWatchedContentQueryParamType = 10
)

func (e GetTopWatchedContentQueryParamType) ToPointer() *GetTopWatchedContentQueryParamType {
	return &e
}

type GetTopWatchedContentRequest struct {
	// Adds the Guids object to the response
	//
	IncludeGuids *int64 `queryParam:"style=form,explode=true,name=includeGuids"`
	// The type of media to retrieve or filter by.
	// 1 = movie
	// 2 = show
	// 3 = season
	// 4 = episode
	// E.g. A movie library will not return anything with type 3 as there are no seasons for movie libraries
	//
	Type GetTopWatchedContentQueryParamType `queryParam:"style=form,explode=true,name=type"`
}

func (o *GetTopWatchedContentRequest) GetIncludeGuids() *int64 {
	if o == nil {
		return nil
	}
	return o.IncludeGuids
}

func (o *GetTopWatchedContentRequest) GetType() GetTopWatchedContentQueryParamType {
	if o == nil {
		return GetTopWatchedContentQueryParamType(0)
	}
	return o.Type
}

type GetTopWatchedContentGenre struct {
	ID     *int    `json:"id,omitempty"`
	Filter *string `json:"filter,omitempty"`
	Tag    *string `json:"tag,omitempty"`
}

func (o *GetTopWatchedContentGenre) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetTopWatchedContentGenre) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetTopWatchedContentGenre) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetTopWatchedContentCountry struct {
	ID     *int    `json:"id,omitempty"`
	Filter *string `json:"filter,omitempty"`
	Tag    *string `json:"tag,omitempty"`
}

func (o *GetTopWatchedContentCountry) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetTopWatchedContentCountry) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetTopWatchedContentCountry) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetTopWatchedContentGuids struct {
	ID *string `json:"id,omitempty"`
}

func (o *GetTopWatchedContentGuids) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type GetTopWatchedContentRole struct {
	ID     *int    `json:"id,omitempty"`
	Filter *string `json:"filter,omitempty"`
	Tag    *string `json:"tag,omitempty"`
	TagKey *string `json:"tagKey,omitempty"`
	Role   *string `json:"role,omitempty"`
	Thumb  *string `json:"thumb,omitempty"`
}

func (o *GetTopWatchedContentRole) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetTopWatchedContentRole) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetTopWatchedContentRole) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

func (o *GetTopWatchedContentRole) GetTagKey() *string {
	if o == nil {
		return nil
	}
	return o.TagKey
}

func (o *GetTopWatchedContentRole) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *GetTopWatchedContentRole) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

type GetTopWatchedContentUser struct {
	ID *int `json:"id,omitempty"`
}

func (o *GetTopWatchedContentUser) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

type GetTopWatchedContentMetadata struct {
	RatingKey             *string                       `json:"ratingKey,omitempty"`
	Key                   *string                       `json:"key,omitempty"`
	GUID                  *string                       `json:"guid,omitempty"`
	Slug                  *string                       `json:"slug,omitempty"`
	Studio                *string                       `json:"studio,omitempty"`
	Type                  *string                       `json:"type,omitempty"`
	Title                 *string                       `json:"title,omitempty"`
	LibrarySectionTitle   *string                       `json:"librarySectionTitle,omitempty"`
	LibrarySectionID      *int                          `json:"librarySectionID,omitempty"`
	LibrarySectionKey     *string                       `json:"librarySectionKey,omitempty"`
	ContentRating         *string                       `json:"contentRating,omitempty"`
	Summary               *string                       `json:"summary,omitempty"`
	Index                 *int64                        `json:"index,omitempty"`
	AudienceRating        *float64                      `json:"audienceRating,omitempty"`
	Year                  *int                          `json:"year,omitempty"`
	Tagline               *string                       `json:"tagline,omitempty"`
	Thumb                 *string                       `json:"thumb,omitempty"`
	Art                   *string                       `json:"art,omitempty"`
	Duration              *int                          `json:"duration,omitempty"`
	OriginallyAvailableAt *types.Date                   `json:"originallyAvailableAt,omitempty"`
	LeafCount             *int64                        `json:"leafCount,omitempty"`
	ViewedLeafCount       *int64                        `json:"viewedLeafCount,omitempty"`
	ChildCount            *int64                        `json:"childCount,omitempty"`
	AddedAt               *int                          `json:"addedAt,omitempty"`
	UpdatedAt             *int                          `json:"updatedAt,omitempty"`
	GlobalViewCount       *int64                        `json:"globalViewCount,omitempty"`
	AudienceRatingImage   *string                       `json:"audienceRatingImage,omitempty"`
	Genre                 []GetTopWatchedContentGenre   `json:"Genre,omitempty"`
	Country               []GetTopWatchedContentCountry `json:"Country,omitempty"`
	Guids                 []GetTopWatchedContentGuids   `json:"Guid,omitempty"`
	Role                  []GetTopWatchedContentRole    `json:"Role,omitempty"`
	User                  []GetTopWatchedContentUser    `json:"User,omitempty"`
}

func (g GetTopWatchedContentMetadata) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTopWatchedContentMetadata) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTopWatchedContentMetadata) GetRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.RatingKey
}

func (o *GetTopWatchedContentMetadata) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetTopWatchedContentMetadata) GetGUID() *string {
	if o == nil {
		return nil
	}
	return o.GUID
}

func (o *GetTopWatchedContentMetadata) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *GetTopWatchedContentMetadata) GetStudio() *string {
	if o == nil {
		return nil
	}
	return o.Studio
}

func (o *GetTopWatchedContentMetadata) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetTopWatchedContentMetadata) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetTopWatchedContentMetadata) GetLibrarySectionTitle() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionTitle
}

func (o *GetTopWatchedContentMetadata) GetLibrarySectionID() *int {
	if o == nil {
		return nil
	}
	return o.LibrarySectionID
}

func (o *GetTopWatchedContentMetadata) GetLibrarySectionKey() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionKey
}

func (o *GetTopWatchedContentMetadata) GetContentRating() *string {
	if o == nil {
		return nil
	}
	return o.ContentRating
}

func (o *GetTopWatchedContentMetadata) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *GetTopWatchedContentMetadata) GetIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *GetTopWatchedContentMetadata) GetAudienceRating() *float64 {
	if o == nil {
		return nil
	}
	return o.AudienceRating
}

func (o *GetTopWatchedContentMetadata) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

func (o *GetTopWatchedContentMetadata) GetTagline() *string {
	if o == nil {
		return nil
	}
	return o.Tagline
}

func (o *GetTopWatchedContentMetadata) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetTopWatchedContentMetadata) GetArt() *string {
	if o == nil {
		return nil
	}
	return o.Art
}

func (o *GetTopWatchedContentMetadata) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetTopWatchedContentMetadata) GetOriginallyAvailableAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.OriginallyAvailableAt
}

func (o *GetTopWatchedContentMetadata) GetLeafCount() *int64 {
	if o == nil {
		return nil
	}
	return o.LeafCount
}

func (o *GetTopWatchedContentMetadata) GetViewedLeafCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ViewedLeafCount
}

func (o *GetTopWatchedContentMetadata) GetChildCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ChildCount
}

func (o *GetTopWatchedContentMetadata) GetAddedAt() *int {
	if o == nil {
		return nil
	}
	return o.AddedAt
}

func (o *GetTopWatchedContentMetadata) GetUpdatedAt() *int {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetTopWatchedContentMetadata) GetGlobalViewCount() *int64 {
	if o == nil {
		return nil
	}
	return o.GlobalViewCount
}

func (o *GetTopWatchedContentMetadata) GetAudienceRatingImage() *string {
	if o == nil {
		return nil
	}
	return o.AudienceRatingImage
}

func (o *GetTopWatchedContentMetadata) GetGenre() []GetTopWatchedContentGenre {
	if o == nil {
		return nil
	}
	return o.Genre
}

func (o *GetTopWatchedContentMetadata) GetCountry() []GetTopWatchedContentCountry {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *GetTopWatchedContentMetadata) GetGuids() []GetTopWatchedContentGuids {
	if o == nil {
		return nil
	}
	return o.Guids
}

func (o *GetTopWatchedContentMetadata) GetRole() []GetTopWatchedContentRole {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *GetTopWatchedContentMetadata) GetUser() []GetTopWatchedContentUser {
	if o == nil {
		return nil
	}
	return o.User
}

type GetTopWatchedContentMediaContainer struct {
	Size            *int                           `json:"size,omitempty"`
	AllowSync       *bool                          `json:"allowSync,omitempty"`
	Identifier      *string                        `json:"identifier,omitempty"`
	MediaTagPrefix  *string                        `json:"mediaTagPrefix,omitempty"`
	MediaTagVersion *int                           `json:"mediaTagVersion,omitempty"`
	Metadata        []GetTopWatchedContentMetadata `json:"Metadata,omitempty"`
}

func (o *GetTopWatchedContentMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetTopWatchedContentMediaContainer) GetAllowSync() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSync
}

func (o *GetTopWatchedContentMediaContainer) GetIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *GetTopWatchedContentMediaContainer) GetMediaTagPrefix() *string {
	if o == nil {
		return nil
	}
	return o.MediaTagPrefix
}

func (o *GetTopWatchedContentMediaContainer) GetMediaTagVersion() *int {
	if o == nil {
		return nil
	}
	return o.MediaTagVersion
}

func (o *GetTopWatchedContentMediaContainer) GetMetadata() []GetTopWatchedContentMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// GetTopWatchedContentResponseBody - The metadata of the library item.
type GetTopWatchedContentResponseBody struct {
	MediaContainer *GetTopWatchedContentMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetTopWatchedContentResponseBody) GetMediaContainer() *GetTopWatchedContentMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetTopWatchedContentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The metadata of the library item.
	Object *GetTopWatchedContentResponseBody
}

func (o *GetTopWatchedContentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTopWatchedContentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTopWatchedContentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTopWatchedContentResponse) GetObject() *GetTopWatchedContentResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
