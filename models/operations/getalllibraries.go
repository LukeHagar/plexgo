// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/LukeHagar/plexgo/internal/utils"
	"net/http"
)

// GetAllLibrariesType - The library type
type GetAllLibrariesType string

const (
	GetAllLibrariesTypeMovie   GetAllLibrariesType = "movie"
	GetAllLibrariesTypeTvShow  GetAllLibrariesType = "show"
	GetAllLibrariesTypeSeason  GetAllLibrariesType = "season"
	GetAllLibrariesTypeEpisode GetAllLibrariesType = "episode"
	GetAllLibrariesTypeArtist  GetAllLibrariesType = "artist"
	GetAllLibrariesTypeAlbum   GetAllLibrariesType = "album"
)

func (e GetAllLibrariesType) ToPointer() *GetAllLibrariesType {
	return &e
}
func (e *GetAllLibrariesType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "movie":
		fallthrough
	case "show":
		fallthrough
	case "season":
		fallthrough
	case "episode":
		fallthrough
	case "artist":
		fallthrough
	case "album":
		*e = GetAllLibrariesType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllLibrariesType: %v", v)
	}
}

// Hidden - UNKNOWN
type Hidden int

const (
	HiddenDisable Hidden = 0
	HiddenEnable  Hidden = 1
)

func (e Hidden) ToPointer() *Hidden {
	return &e
}
func (e *Hidden) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = Hidden(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Hidden: %v", v)
	}
}

type GetAllLibrariesLocation struct {
	// The ID of the location.
	ID int `json:"id"`
	// The path to the media item.
	Path string `json:"path"`
}

func (o *GetAllLibrariesLocation) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetAllLibrariesLocation) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

type GetAllLibrariesDirectory struct {
	// Indicates whether syncing is allowed.
	AllowSync bool `json:"allowSync"`
	// URL for the background artwork of the media container.
	Art string `json:"art"`
	// The relative path to the composite media item.
	Composite string `json:"composite"`
	// UNKNOWN
	Filters bool `json:"filters"`
	// Indicates whether the library is currently being refreshed or updated
	Refreshing bool `json:"refreshing"`
	// URL for the thumbnail image of the media container.
	Thumb string `json:"thumb"`
	// The library key representing the unique identifier
	Key  string              `json:"key"`
	Type GetAllLibrariesType `json:"type"`
	// The title of the library
	Title string `json:"title"`
	// The Plex agent used to match and retrieve media metadata.
	Agent string `json:"agent"`
	// UNKNOWN
	Scanner string `json:"scanner"`
	// The Plex library language that has been set
	Language string `json:"language"`
	// The universally unique identifier for the library.
	UUID string `json:"uuid"`
	// Unix epoch datetime in seconds
	UpdatedAt int64  `json:"updatedAt"`
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// Unix epoch datetime in seconds
	ScannedAt int64 `json:"scannedAt"`
	// UNKNOWN
	Content bool `json:"content"`
	// UNKNOWN
	Directory bool `json:"directory"`
	// The number of seconds since the content was last changed relative to now.
	ContentChangedAt int                       `json:"contentChangedAt"`
	Hidden           *Hidden                   `default:"0" json:"hidden"`
	Location         []GetAllLibrariesLocation `json:"Location"`
}

func (g GetAllLibrariesDirectory) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetAllLibrariesDirectory) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetAllLibrariesDirectory) GetAllowSync() bool {
	if o == nil {
		return false
	}
	return o.AllowSync
}

func (o *GetAllLibrariesDirectory) GetArt() string {
	if o == nil {
		return ""
	}
	return o.Art
}

func (o *GetAllLibrariesDirectory) GetComposite() string {
	if o == nil {
		return ""
	}
	return o.Composite
}

func (o *GetAllLibrariesDirectory) GetFilters() bool {
	if o == nil {
		return false
	}
	return o.Filters
}

func (o *GetAllLibrariesDirectory) GetRefreshing() bool {
	if o == nil {
		return false
	}
	return o.Refreshing
}

func (o *GetAllLibrariesDirectory) GetThumb() string {
	if o == nil {
		return ""
	}
	return o.Thumb
}

func (o *GetAllLibrariesDirectory) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *GetAllLibrariesDirectory) GetType() GetAllLibrariesType {
	if o == nil {
		return GetAllLibrariesType("")
	}
	return o.Type
}

func (o *GetAllLibrariesDirectory) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *GetAllLibrariesDirectory) GetAgent() string {
	if o == nil {
		return ""
	}
	return o.Agent
}

func (o *GetAllLibrariesDirectory) GetScanner() string {
	if o == nil {
		return ""
	}
	return o.Scanner
}

func (o *GetAllLibrariesDirectory) GetLanguage() string {
	if o == nil {
		return ""
	}
	return o.Language
}

func (o *GetAllLibrariesDirectory) GetUUID() string {
	if o == nil {
		return ""
	}
	return o.UUID
}

func (o *GetAllLibrariesDirectory) GetUpdatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.UpdatedAt
}

func (o *GetAllLibrariesDirectory) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetAllLibrariesDirectory) GetScannedAt() int64 {
	if o == nil {
		return 0
	}
	return o.ScannedAt
}

func (o *GetAllLibrariesDirectory) GetContent() bool {
	if o == nil {
		return false
	}
	return o.Content
}

func (o *GetAllLibrariesDirectory) GetDirectory() bool {
	if o == nil {
		return false
	}
	return o.Directory
}

func (o *GetAllLibrariesDirectory) GetContentChangedAt() int {
	if o == nil {
		return 0
	}
	return o.ContentChangedAt
}

func (o *GetAllLibrariesDirectory) GetHidden() *Hidden {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *GetAllLibrariesDirectory) GetLocation() []GetAllLibrariesLocation {
	if o == nil {
		return []GetAllLibrariesLocation{}
	}
	return o.Location
}

type GetAllLibrariesMediaContainer struct {
	// Number of media items returned in this response.
	Size int `json:"size"`
	// Indicates whether syncing is allowed.
	AllowSync bool `json:"allowSync"`
	// The primary title of the media container.
	Title1    string                     `json:"title1"`
	Directory []GetAllLibrariesDirectory `json:"Directory,omitempty"`
}

func (o *GetAllLibrariesMediaContainer) GetSize() int {
	if o == nil {
		return 0
	}
	return o.Size
}

func (o *GetAllLibrariesMediaContainer) GetAllowSync() bool {
	if o == nil {
		return false
	}
	return o.AllowSync
}

func (o *GetAllLibrariesMediaContainer) GetTitle1() string {
	if o == nil {
		return ""
	}
	return o.Title1
}

func (o *GetAllLibrariesMediaContainer) GetDirectory() []GetAllLibrariesDirectory {
	if o == nil {
		return nil
	}
	return o.Directory
}

// GetAllLibrariesResponseBody - The libraries available on the Server
type GetAllLibrariesResponseBody struct {
	MediaContainer *GetAllLibrariesMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetAllLibrariesResponseBody) GetMediaContainer() *GetAllLibrariesMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetAllLibrariesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The libraries available on the Server
	Object *GetAllLibrariesResponseBody
}

func (o *GetAllLibrariesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllLibrariesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllLibrariesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllLibrariesResponse) GetObject() *GetAllLibrariesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
