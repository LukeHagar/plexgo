// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetStatisticsRequest struct {
	// The timespan to retrieve statistics for
	// the exact meaning of this parameter is not known
	//
	Timespan *int64 `queryParam:"style=form,explode=true,name=timespan"`
}

func (o *GetStatisticsRequest) GetTimespan() *int64 {
	if o == nil {
		return nil
	}
	return o.Timespan
}

type GetStatisticsDevice struct {
	ID               *int    `json:"id,omitempty"`
	Name             *string `json:"name,omitempty"`
	Platform         *string `json:"platform,omitempty"`
	ClientIdentifier *string `json:"clientIdentifier,omitempty"`
	CreatedAt        *int    `json:"createdAt,omitempty"`
}

func (o *GetStatisticsDevice) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetStatisticsDevice) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetStatisticsDevice) GetPlatform() *string {
	if o == nil {
		return nil
	}
	return o.Platform
}

func (o *GetStatisticsDevice) GetClientIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.ClientIdentifier
}

func (o *GetStatisticsDevice) GetCreatedAt() *int {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

type Account struct {
	ID                      *int    `json:"id,omitempty"`
	Key                     *string `json:"key,omitempty"`
	Name                    *string `json:"name,omitempty"`
	DefaultAudioLanguage    *string `json:"defaultAudioLanguage,omitempty"`
	AutoSelectAudio         *bool   `json:"autoSelectAudio,omitempty"`
	DefaultSubtitleLanguage *string `json:"defaultSubtitleLanguage,omitempty"`
	SubtitleMode            *int    `json:"subtitleMode,omitempty"`
	Thumb                   *string `json:"thumb,omitempty"`
}

func (o *Account) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Account) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *Account) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Account) GetDefaultAudioLanguage() *string {
	if o == nil {
		return nil
	}
	return o.DefaultAudioLanguage
}

func (o *Account) GetAutoSelectAudio() *bool {
	if o == nil {
		return nil
	}
	return o.AutoSelectAudio
}

func (o *Account) GetDefaultSubtitleLanguage() *string {
	if o == nil {
		return nil
	}
	return o.DefaultSubtitleLanguage
}

func (o *Account) GetSubtitleMode() *int {
	if o == nil {
		return nil
	}
	return o.SubtitleMode
}

func (o *Account) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

type StatisticsMedia struct {
	AccountID    *int `json:"accountID,omitempty"`
	DeviceID     *int `json:"deviceID,omitempty"`
	Timespan     *int `json:"timespan,omitempty"`
	At           *int `json:"at,omitempty"`
	MetadataType *int `json:"metadataType,omitempty"`
	Count        *int `json:"count,omitempty"`
	Duration     *int `json:"duration,omitempty"`
}

func (o *StatisticsMedia) GetAccountID() *int {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *StatisticsMedia) GetDeviceID() *int {
	if o == nil {
		return nil
	}
	return o.DeviceID
}

func (o *StatisticsMedia) GetTimespan() *int {
	if o == nil {
		return nil
	}
	return o.Timespan
}

func (o *StatisticsMedia) GetAt() *int {
	if o == nil {
		return nil
	}
	return o.At
}

func (o *StatisticsMedia) GetMetadataType() *int {
	if o == nil {
		return nil
	}
	return o.MetadataType
}

func (o *StatisticsMedia) GetCount() *int {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *StatisticsMedia) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

type GetStatisticsMediaContainer struct {
	Size            *int                  `json:"size,omitempty"`
	Device          []GetStatisticsDevice `json:"Device,omitempty"`
	Account         []Account             `json:"Account,omitempty"`
	StatisticsMedia []StatisticsMedia     `json:"StatisticsMedia,omitempty"`
}

func (o *GetStatisticsMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetStatisticsMediaContainer) GetDevice() []GetStatisticsDevice {
	if o == nil {
		return nil
	}
	return o.Device
}

func (o *GetStatisticsMediaContainer) GetAccount() []Account {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *GetStatisticsMediaContainer) GetStatisticsMedia() []StatisticsMedia {
	if o == nil {
		return nil
	}
	return o.StatisticsMedia
}

// GetStatisticsResponseBody - Media Statistics
type GetStatisticsResponseBody struct {
	MediaContainer *GetStatisticsMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetStatisticsResponseBody) GetMediaContainer() *GetStatisticsMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetStatisticsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Media Statistics
	Object *GetStatisticsResponseBody
}

func (o *GetStatisticsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetStatisticsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetStatisticsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetStatisticsResponse) GetObject() *GetStatisticsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
