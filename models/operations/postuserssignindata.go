// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/LukeHagar/plexgo/internal/utils"
	"net/http"
)

var PostUsersSignInDataServerList = []string{
	"https://plex.tv/api/v2",
}

// PostUsersSignInDataRequestBody - Login credentials
type PostUsersSignInDataRequestBody struct {
	Login            string  `form:"name=login"`
	Password         string  `form:"name=password"`
	RememberMe       *bool   `default:"false" form:"name=rememberMe"`
	VerificationCode *string `form:"name=verificationCode"`
}

func (p PostUsersSignInDataRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostUsersSignInDataRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostUsersSignInDataRequestBody) GetLogin() string {
	if o == nil {
		return ""
	}
	return o.Login
}

func (o *PostUsersSignInDataRequestBody) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *PostUsersSignInDataRequestBody) GetRememberMe() *bool {
	if o == nil {
		return nil
	}
	return o.RememberMe
}

func (o *PostUsersSignInDataRequestBody) GetVerificationCode() *string {
	if o == nil {
		return nil
	}
	return o.VerificationCode
}

type PostUsersSignInDataRequest struct {
	// An opaque identifier unique to the client (UUID, serial number, or other unique device ID)
	ClientID string `header:"style=simple,explode=false,name=X-Plex-Client-Identifier"`
	// The name of the client application. (Plex Web, Plex Media Server, etc.)
	ClientName *string `header:"style=simple,explode=false,name=X-Plex-Product"`
	// A relatively friendly name for the client device
	DeviceNickname *string `header:"style=simple,explode=false,name=X-Plex-Device"`
	// The version of the client application.
	ClientVersion *string `header:"style=simple,explode=false,name=X-Plex-Version"`
	// The platform of the client application.
	Platform *string `header:"style=simple,explode=false,name=X-Plex-Platform"`
	// Login credentials
	RequestBody *PostUsersSignInDataRequestBody `request:"mediaType=application/x-www-form-urlencoded"`
}

func (o *PostUsersSignInDataRequest) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *PostUsersSignInDataRequest) GetClientName() *string {
	if o == nil {
		return nil
	}
	return o.ClientName
}

func (o *PostUsersSignInDataRequest) GetDeviceNickname() *string {
	if o == nil {
		return nil
	}
	return o.DeviceNickname
}

func (o *PostUsersSignInDataRequest) GetClientVersion() *string {
	if o == nil {
		return nil
	}
	return o.ClientVersion
}

func (o *PostUsersSignInDataRequest) GetPlatform() *string {
	if o == nil {
		return nil
	}
	return o.Platform
}

func (o *PostUsersSignInDataRequest) GetRequestBody() *PostUsersSignInDataRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

// PostUsersSignInDataMailingListStatus - Your current mailing list status (active or unsubscribed)
type PostUsersSignInDataMailingListStatus string

const (
	PostUsersSignInDataMailingListStatusActive       PostUsersSignInDataMailingListStatus = "active"
	PostUsersSignInDataMailingListStatusUnsubscribed PostUsersSignInDataMailingListStatus = "unsubscribed"
)

func (e PostUsersSignInDataMailingListStatus) ToPointer() *PostUsersSignInDataMailingListStatus {
	return &e
}
func (e *PostUsersSignInDataMailingListStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "unsubscribed":
		*e = PostUsersSignInDataMailingListStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostUsersSignInDataMailingListStatus: %v", v)
	}
}

// PostUsersSignInDataAutoSelectSubtitle - The auto-select subtitle mode (0 = Manually selected, 1 = Shown with foreign audio, 2 = Always enabled)
type PostUsersSignInDataAutoSelectSubtitle int

const (
	PostUsersSignInDataAutoSelectSubtitleDisable PostUsersSignInDataAutoSelectSubtitle = 0
	PostUsersSignInDataAutoSelectSubtitleEnable  PostUsersSignInDataAutoSelectSubtitle = 1
)

func (e PostUsersSignInDataAutoSelectSubtitle) ToPointer() *PostUsersSignInDataAutoSelectSubtitle {
	return &e
}
func (e *PostUsersSignInDataAutoSelectSubtitle) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = PostUsersSignInDataAutoSelectSubtitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostUsersSignInDataAutoSelectSubtitle: %v", v)
	}
}

// PostUsersSignInDataDefaultSubtitleAccessibility - The subtitles for the deaf or hard-of-hearing (SDH) searches mode (0 = Prefer non-SDH subtitles, 1 = Prefer SDH subtitles, 2 = Only show SDH subtitles, 3 = Only show non-SDH subtitles)
type PostUsersSignInDataDefaultSubtitleAccessibility int

const (
	PostUsersSignInDataDefaultSubtitleAccessibilityDisable PostUsersSignInDataDefaultSubtitleAccessibility = 0
	PostUsersSignInDataDefaultSubtitleAccessibilityEnable  PostUsersSignInDataDefaultSubtitleAccessibility = 1
)

func (e PostUsersSignInDataDefaultSubtitleAccessibility) ToPointer() *PostUsersSignInDataDefaultSubtitleAccessibility {
	return &e
}
func (e *PostUsersSignInDataDefaultSubtitleAccessibility) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = PostUsersSignInDataDefaultSubtitleAccessibility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostUsersSignInDataDefaultSubtitleAccessibility: %v", v)
	}
}

// PostUsersSignInDataDefaultSubtitleForced - The forced subtitles searches mode (0 = Prefer non-forced subtitles, 1 = Prefer forced subtitles, 2 = Only show forced subtitles, 3 = Only show non-forced subtitles)
type PostUsersSignInDataDefaultSubtitleForced int

const (
	PostUsersSignInDataDefaultSubtitleForcedDisable PostUsersSignInDataDefaultSubtitleForced = 0
	PostUsersSignInDataDefaultSubtitleForcedEnable  PostUsersSignInDataDefaultSubtitleForced = 1
)

func (e PostUsersSignInDataDefaultSubtitleForced) ToPointer() *PostUsersSignInDataDefaultSubtitleForced {
	return &e
}
func (e *PostUsersSignInDataDefaultSubtitleForced) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = PostUsersSignInDataDefaultSubtitleForced(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostUsersSignInDataDefaultSubtitleForced: %v", v)
	}
}

// PostUsersSignInDataWatchedIndicator - Whether or not media watched indicators are enabled (little orange dot on media)
type PostUsersSignInDataWatchedIndicator int

const (
	PostUsersSignInDataWatchedIndicatorDisable PostUsersSignInDataWatchedIndicator = 0
	PostUsersSignInDataWatchedIndicatorEnable  PostUsersSignInDataWatchedIndicator = 1
)

func (e PostUsersSignInDataWatchedIndicator) ToPointer() *PostUsersSignInDataWatchedIndicator {
	return &e
}
func (e *PostUsersSignInDataWatchedIndicator) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = PostUsersSignInDataWatchedIndicator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostUsersSignInDataWatchedIndicator: %v", v)
	}
}

// PostUsersSignInDataMediaReviewsVisibility - Whether or not the account has media reviews visibility enabled
type PostUsersSignInDataMediaReviewsVisibility int

const (
	PostUsersSignInDataMediaReviewsVisibilityDisable PostUsersSignInDataMediaReviewsVisibility = 0
	PostUsersSignInDataMediaReviewsVisibilityEnable  PostUsersSignInDataMediaReviewsVisibility = 1
)

func (e PostUsersSignInDataMediaReviewsVisibility) ToPointer() *PostUsersSignInDataMediaReviewsVisibility {
	return &e
}
func (e *PostUsersSignInDataMediaReviewsVisibility) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = PostUsersSignInDataMediaReviewsVisibility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostUsersSignInDataMediaReviewsVisibility: %v", v)
	}
}

type PostUsersSignInDataUserProfile struct {
	// If the account has automatically select audio and subtitle tracks enabled
	AutoSelectAudio *bool `default:"true" json:"autoSelectAudio"`
	// The preferred audio language for the account
	DefaultAudioLanguage *string `json:"defaultAudioLanguage"`
	// The preferred subtitle language for the account
	DefaultSubtitleLanguage      *string                                          `json:"defaultSubtitleLanguage"`
	AutoSelectSubtitle           *PostUsersSignInDataAutoSelectSubtitle           `default:"0" json:"autoSelectSubtitle"`
	DefaultSubtitleAccessibility *PostUsersSignInDataDefaultSubtitleAccessibility `default:"0" json:"defaultSubtitleAccessibility"`
	DefaultSubtitleForced        *PostUsersSignInDataDefaultSubtitleForced        `default:"0" json:"defaultSubtitleForced"`
	WatchedIndicator             *PostUsersSignInDataWatchedIndicator             `default:"0" json:"watchedIndicator"`
	MediaReviewsVisibility       *PostUsersSignInDataMediaReviewsVisibility       `default:"0" json:"mediaReviewsVisibility"`
}

func (p PostUsersSignInDataUserProfile) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostUsersSignInDataUserProfile) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostUsersSignInDataUserProfile) GetAutoSelectAudio() *bool {
	if o == nil {
		return nil
	}
	return o.AutoSelectAudio
}

func (o *PostUsersSignInDataUserProfile) GetDefaultAudioLanguage() *string {
	if o == nil {
		return nil
	}
	return o.DefaultAudioLanguage
}

func (o *PostUsersSignInDataUserProfile) GetDefaultSubtitleLanguage() *string {
	if o == nil {
		return nil
	}
	return o.DefaultSubtitleLanguage
}

func (o *PostUsersSignInDataUserProfile) GetAutoSelectSubtitle() *PostUsersSignInDataAutoSelectSubtitle {
	if o == nil {
		return nil
	}
	return o.AutoSelectSubtitle
}

func (o *PostUsersSignInDataUserProfile) GetDefaultSubtitleAccessibility() *PostUsersSignInDataDefaultSubtitleAccessibility {
	if o == nil {
		return nil
	}
	return o.DefaultSubtitleAccessibility
}

func (o *PostUsersSignInDataUserProfile) GetDefaultSubtitleForced() *PostUsersSignInDataDefaultSubtitleForced {
	if o == nil {
		return nil
	}
	return o.DefaultSubtitleForced
}

func (o *PostUsersSignInDataUserProfile) GetWatchedIndicator() *PostUsersSignInDataWatchedIndicator {
	if o == nil {
		return nil
	}
	return o.WatchedIndicator
}

func (o *PostUsersSignInDataUserProfile) GetMediaReviewsVisibility() *PostUsersSignInDataMediaReviewsVisibility {
	if o == nil {
		return nil
	}
	return o.MediaReviewsVisibility
}

type PostUsersSignInDataStatus string

const (
	PostUsersSignInDataStatusOnline  PostUsersSignInDataStatus = "online"
	PostUsersSignInDataStatusOffline PostUsersSignInDataStatus = "offline"
)

func (e PostUsersSignInDataStatus) ToPointer() *PostUsersSignInDataStatus {
	return &e
}
func (e *PostUsersSignInDataStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "online":
		fallthrough
	case "offline":
		*e = PostUsersSignInDataStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostUsersSignInDataStatus: %v", v)
	}
}

type PostUsersSignInDataServices struct {
	Identifier string                    `json:"identifier"`
	Endpoint   string                    `json:"endpoint"`
	Token      *string                   `json:"token"`
	Secret     *string                   `json:"secret"`
	Status     PostUsersSignInDataStatus `json:"status"`
}

func (o *PostUsersSignInDataServices) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *PostUsersSignInDataServices) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *PostUsersSignInDataServices) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *PostUsersSignInDataServices) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *PostUsersSignInDataServices) GetStatus() PostUsersSignInDataStatus {
	if o == nil {
		return PostUsersSignInDataStatus("")
	}
	return o.Status
}

// PostUsersSignInDataAuthenticationStatus - String representation of subscriptionActive
type PostUsersSignInDataAuthenticationStatus string

const (
	PostUsersSignInDataAuthenticationStatusInactive PostUsersSignInDataAuthenticationStatus = "Inactive"
	PostUsersSignInDataAuthenticationStatusActive   PostUsersSignInDataAuthenticationStatus = "Active"
)

func (e PostUsersSignInDataAuthenticationStatus) ToPointer() *PostUsersSignInDataAuthenticationStatus {
	return &e
}
func (e *PostUsersSignInDataAuthenticationStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Inactive":
		fallthrough
	case "Active":
		*e = PostUsersSignInDataAuthenticationStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostUsersSignInDataAuthenticationStatus: %v", v)
	}
}

// PostUsersSignInDataSubscription - If the account’s Plex Pass subscription is active
type PostUsersSignInDataSubscription struct {
	// List of features allowed on your Plex Pass subscription
	Features []string `json:"features,omitempty"`
	// If the account's Plex Pass subscription is active
	Active *bool `json:"active,omitempty"`
	// Date the account subscribed to Plex Pass
	SubscribedAt *string `json:"subscribedAt,omitempty"`
	// String representation of subscriptionActive
	Status *PostUsersSignInDataAuthenticationStatus `json:"status,omitempty"`
	// Payment service used for your Plex Pass subscription
	PaymentService *string `json:"paymentService,omitempty"`
	// Name of Plex Pass subscription plan
	Plan *string `json:"plan,omitempty"`
}

func (o *PostUsersSignInDataSubscription) GetFeatures() []string {
	if o == nil {
		return nil
	}
	return o.Features
}

func (o *PostUsersSignInDataSubscription) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *PostUsersSignInDataSubscription) GetSubscribedAt() *string {
	if o == nil {
		return nil
	}
	return o.SubscribedAt
}

func (o *PostUsersSignInDataSubscription) GetStatus() *PostUsersSignInDataAuthenticationStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *PostUsersSignInDataSubscription) GetPaymentService() *string {
	if o == nil {
		return nil
	}
	return o.PaymentService
}

func (o *PostUsersSignInDataSubscription) GetPlan() *string {
	if o == nil {
		return nil
	}
	return o.Plan
}

// PostUsersSignInDataAuthenticationResponseStatus - String representation of subscriptionActive
type PostUsersSignInDataAuthenticationResponseStatus string

const (
	PostUsersSignInDataAuthenticationResponseStatusInactive PostUsersSignInDataAuthenticationResponseStatus = "Inactive"
	PostUsersSignInDataAuthenticationResponseStatusActive   PostUsersSignInDataAuthenticationResponseStatus = "Active"
)

func (e PostUsersSignInDataAuthenticationResponseStatus) ToPointer() *PostUsersSignInDataAuthenticationResponseStatus {
	return &e
}
func (e *PostUsersSignInDataAuthenticationResponseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Inactive":
		fallthrough
	case "Active":
		*e = PostUsersSignInDataAuthenticationResponseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostUsersSignInDataAuthenticationResponseStatus: %v", v)
	}
}

type PostUsersSignInDataAuthenticationSubscription struct {
	// List of features allowed on your Plex Pass subscription
	Features []string `json:"features,omitempty"`
	// If the account's Plex Pass subscription is active
	Active *bool `json:"active,omitempty"`
	// Date the account subscribed to Plex Pass
	SubscribedAt *string `json:"subscribedAt,omitempty"`
	// String representation of subscriptionActive
	Status *PostUsersSignInDataAuthenticationResponseStatus `json:"status,omitempty"`
	// Payment service used for your Plex Pass subscription
	PaymentService *string `json:"paymentService,omitempty"`
	// Name of Plex Pass subscription plan
	Plan *string `json:"plan,omitempty"`
}

func (o *PostUsersSignInDataAuthenticationSubscription) GetFeatures() []string {
	if o == nil {
		return nil
	}
	return o.Features
}

func (o *PostUsersSignInDataAuthenticationSubscription) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *PostUsersSignInDataAuthenticationSubscription) GetSubscribedAt() *string {
	if o == nil {
		return nil
	}
	return o.SubscribedAt
}

func (o *PostUsersSignInDataAuthenticationSubscription) GetStatus() *PostUsersSignInDataAuthenticationResponseStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *PostUsersSignInDataAuthenticationSubscription) GetPaymentService() *string {
	if o == nil {
		return nil
	}
	return o.PaymentService
}

func (o *PostUsersSignInDataAuthenticationSubscription) GetPlan() *string {
	if o == nil {
		return nil
	}
	return o.Plan
}

type PostUsersSignInDataState string

const (
	PostUsersSignInDataStateEnded PostUsersSignInDataState = "ended"
)

func (e PostUsersSignInDataState) ToPointer() *PostUsersSignInDataState {
	return &e
}
func (e *PostUsersSignInDataState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ended":
		*e = PostUsersSignInDataState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostUsersSignInDataState: %v", v)
	}
}

type InternalPaymentMethod struct {
}

type Billing struct {
	InternalPaymentMethod InternalPaymentMethod `json:"internalPaymentMethod"`
	PaymentMethodID       *int64                `json:"paymentMethodId"`
}

func (o *Billing) GetInternalPaymentMethod() InternalPaymentMethod {
	if o == nil {
		return InternalPaymentMethod{}
	}
	return o.InternalPaymentMethod
}

func (o *Billing) GetPaymentMethodID() *int64 {
	if o == nil {
		return nil
	}
	return o.PaymentMethodID
}

type PastSubscription struct {
	ID            *string                  `json:"id"`
	Mode          *string                  `json:"mode"`
	RenewsAt      *int64                   `json:"renewsAt"`
	EndsAt        *int64                   `json:"endsAt"`
	Canceled      *bool                    `default:"false" json:"canceled"`
	GracePeriod   *bool                    `default:"false" json:"gracePeriod"`
	OnHold        *bool                    `default:"false" json:"onHold"`
	CanReactivate *bool                    `default:"false" json:"canReactivate"`
	CanUpgrade    *bool                    `default:"false" json:"canUpgrade"`
	CanDowngrade  *bool                    `default:"false" json:"canDowngrade"`
	CanConvert    *bool                    `default:"false" json:"canConvert"`
	Type          string                   `json:"type"`
	Transfer      *string                  `json:"transfer"`
	State         PostUsersSignInDataState `json:"state"`
	Billing       Billing                  `json:"billing"`
}

func (p PastSubscription) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PastSubscription) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PastSubscription) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PastSubscription) GetMode() *string {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *PastSubscription) GetRenewsAt() *int64 {
	if o == nil {
		return nil
	}
	return o.RenewsAt
}

func (o *PastSubscription) GetEndsAt() *int64 {
	if o == nil {
		return nil
	}
	return o.EndsAt
}

func (o *PastSubscription) GetCanceled() *bool {
	if o == nil {
		return nil
	}
	return o.Canceled
}

func (o *PastSubscription) GetGracePeriod() *bool {
	if o == nil {
		return nil
	}
	return o.GracePeriod
}

func (o *PastSubscription) GetOnHold() *bool {
	if o == nil {
		return nil
	}
	return o.OnHold
}

func (o *PastSubscription) GetCanReactivate() *bool {
	if o == nil {
		return nil
	}
	return o.CanReactivate
}

func (o *PastSubscription) GetCanUpgrade() *bool {
	if o == nil {
		return nil
	}
	return o.CanUpgrade
}

func (o *PastSubscription) GetCanDowngrade() *bool {
	if o == nil {
		return nil
	}
	return o.CanDowngrade
}

func (o *PastSubscription) GetCanConvert() *bool {
	if o == nil {
		return nil
	}
	return o.CanConvert
}

func (o *PastSubscription) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *PastSubscription) GetTransfer() *string {
	if o == nil {
		return nil
	}
	return o.Transfer
}

func (o *PastSubscription) GetState() PostUsersSignInDataState {
	if o == nil {
		return PostUsersSignInDataState("")
	}
	return o.State
}

func (o *PastSubscription) GetBilling() Billing {
	if o == nil {
		return Billing{}
	}
	return o.Billing
}

type Trials struct {
}

// PostUsersSignInDataUserPlexAccount - Returns the user account data with a valid auth token
type PostUsersSignInDataUserPlexAccount struct {
	// Unknown
	AdsConsent           *bool  `json:"adsConsent"`
	AdsConsentReminderAt *int64 `json:"adsConsentReminderAt"`
	AdsConsentSetAt      *int64 `json:"adsConsentSetAt"`
	// Unknown
	Anonymous *bool `default:"false" json:"anonymous"`
	// The account token
	AuthToken string `json:"authToken"`
	// If the two-factor authentication backup codes have been created
	BackupCodesCreated *bool `default:"false" json:"backupCodesCreated"`
	// If the account has been confirmed
	Confirmed *bool `default:"false" json:"confirmed"`
	// The account country
	Country string `json:"country"`
	// The account email address
	Email string `json:"email"`
	// If login with email only is enabled
	EmailOnlyAuth *bool `default:"false" json:"emailOnlyAuth"`
	// If experimental features are enabled
	ExperimentalFeatures *bool `default:"false" json:"experimentalFeatures"`
	// Your account full name
	FriendlyName string `json:"friendlyName"`
	// List of devices your allowed to use with this account
	Entitlements []string `json:"entitlements"`
	// If the account is a Plex Home guest user
	Guest *bool `default:"false" json:"guest"`
	// If the account has a password
	HasPassword *bool `default:"true" json:"hasPassword"`
	// If the account is a Plex Home user
	Home *bool `default:"false" json:"home"`
	// If the account is the Plex Home admin
	HomeAdmin *bool `default:"false" json:"homeAdmin"`
	// The number of accounts in the Plex Home
	HomeSize int `json:"homeSize"`
	// The Plex account ID
	ID int `json:"id"`
	// Unix epoch datetime in seconds
	JoinedAt int64 `json:"joinedAt"`
	// The account locale
	Locale *string `json:"locale"`
	// If you are subscribed to the Plex newsletter
	MailingListActive *bool `default:"false" json:"mailingListActive"`
	// Your current mailing list status (active or unsubscribed)
	MailingListStatus PostUsersSignInDataMailingListStatus `json:"mailingListStatus"`
	// The maximum number of accounts allowed in the Plex Home
	MaxHomeSize int `json:"maxHomeSize"`
	// [Might be removed] The hashed Plex Home PIN
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	Pin     *string                        `json:"pin,omitempty"`
	Profile PostUsersSignInDataUserProfile `json:"profile"`
	// If the account has a Plex Home PIN enabled
	Protected *bool `default:"false" json:"protected"`
	// Unix epoch datetime in seconds
	RememberExpiresAt int64 `json:"rememberExpiresAt"`
	// If the account is a Plex Home managed user
	Restricted *bool `default:"false" json:"restricted"`
	// [Might be removed] List of account roles. Plexpass membership listed here
	Roles []string `json:"roles,omitempty"`
	// Unknown
	ScrobbleTypes string                        `json:"scrobbleTypes"`
	Services      []PostUsersSignInDataServices `json:"services"`
	// If the account’s Plex Pass subscription is active
	Subscription PostUsersSignInDataSubscription `json:"subscription"`
	// Description of the Plex Pass subscription
	SubscriptionDescription *string                                         `json:"subscriptionDescription"`
	Subscriptions           []PostUsersSignInDataAuthenticationSubscription `json:"subscriptions"`
	// URL of the account thumbnail
	Thumb string `json:"thumb"`
	// The title of the account (username or friendly name)
	Title string `json:"title"`
	// If two-factor authentication is enabled
	TwoFactorEnabled *bool `default:"false" json:"twoFactorEnabled"`
	// The account username
	Username string `json:"username"`
	// The account UUID
	UUID               string             `json:"uuid"`
	AttributionPartner *string            `json:"attributionPartner"`
	PastSubscriptions  []PastSubscription `json:"pastSubscriptions"`
	Trials             []Trials           `json:"trials"`
}

func (p PostUsersSignInDataUserPlexAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostUsersSignInDataUserPlexAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostUsersSignInDataUserPlexAccount) GetAdsConsent() *bool {
	if o == nil {
		return nil
	}
	return o.AdsConsent
}

func (o *PostUsersSignInDataUserPlexAccount) GetAdsConsentReminderAt() *int64 {
	if o == nil {
		return nil
	}
	return o.AdsConsentReminderAt
}

func (o *PostUsersSignInDataUserPlexAccount) GetAdsConsentSetAt() *int64 {
	if o == nil {
		return nil
	}
	return o.AdsConsentSetAt
}

func (o *PostUsersSignInDataUserPlexAccount) GetAnonymous() *bool {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *PostUsersSignInDataUserPlexAccount) GetAuthToken() string {
	if o == nil {
		return ""
	}
	return o.AuthToken
}

func (o *PostUsersSignInDataUserPlexAccount) GetBackupCodesCreated() *bool {
	if o == nil {
		return nil
	}
	return o.BackupCodesCreated
}

func (o *PostUsersSignInDataUserPlexAccount) GetConfirmed() *bool {
	if o == nil {
		return nil
	}
	return o.Confirmed
}

func (o *PostUsersSignInDataUserPlexAccount) GetCountry() string {
	if o == nil {
		return ""
	}
	return o.Country
}

func (o *PostUsersSignInDataUserPlexAccount) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *PostUsersSignInDataUserPlexAccount) GetEmailOnlyAuth() *bool {
	if o == nil {
		return nil
	}
	return o.EmailOnlyAuth
}

func (o *PostUsersSignInDataUserPlexAccount) GetExperimentalFeatures() *bool {
	if o == nil {
		return nil
	}
	return o.ExperimentalFeatures
}

func (o *PostUsersSignInDataUserPlexAccount) GetFriendlyName() string {
	if o == nil {
		return ""
	}
	return o.FriendlyName
}

func (o *PostUsersSignInDataUserPlexAccount) GetEntitlements() []string {
	if o == nil {
		return []string{}
	}
	return o.Entitlements
}

func (o *PostUsersSignInDataUserPlexAccount) GetGuest() *bool {
	if o == nil {
		return nil
	}
	return o.Guest
}

func (o *PostUsersSignInDataUserPlexAccount) GetHasPassword() *bool {
	if o == nil {
		return nil
	}
	return o.HasPassword
}

func (o *PostUsersSignInDataUserPlexAccount) GetHome() *bool {
	if o == nil {
		return nil
	}
	return o.Home
}

func (o *PostUsersSignInDataUserPlexAccount) GetHomeAdmin() *bool {
	if o == nil {
		return nil
	}
	return o.HomeAdmin
}

func (o *PostUsersSignInDataUserPlexAccount) GetHomeSize() int {
	if o == nil {
		return 0
	}
	return o.HomeSize
}

func (o *PostUsersSignInDataUserPlexAccount) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *PostUsersSignInDataUserPlexAccount) GetJoinedAt() int64 {
	if o == nil {
		return 0
	}
	return o.JoinedAt
}

func (o *PostUsersSignInDataUserPlexAccount) GetLocale() *string {
	if o == nil {
		return nil
	}
	return o.Locale
}

func (o *PostUsersSignInDataUserPlexAccount) GetMailingListActive() *bool {
	if o == nil {
		return nil
	}
	return o.MailingListActive
}

func (o *PostUsersSignInDataUserPlexAccount) GetMailingListStatus() PostUsersSignInDataMailingListStatus {
	if o == nil {
		return PostUsersSignInDataMailingListStatus("")
	}
	return o.MailingListStatus
}

func (o *PostUsersSignInDataUserPlexAccount) GetMaxHomeSize() int {
	if o == nil {
		return 0
	}
	return o.MaxHomeSize
}

func (o *PostUsersSignInDataUserPlexAccount) GetPin() *string {
	if o == nil {
		return nil
	}
	return o.Pin
}

func (o *PostUsersSignInDataUserPlexAccount) GetProfile() PostUsersSignInDataUserProfile {
	if o == nil {
		return PostUsersSignInDataUserProfile{}
	}
	return o.Profile
}

func (o *PostUsersSignInDataUserPlexAccount) GetProtected() *bool {
	if o == nil {
		return nil
	}
	return o.Protected
}

func (o *PostUsersSignInDataUserPlexAccount) GetRememberExpiresAt() int64 {
	if o == nil {
		return 0
	}
	return o.RememberExpiresAt
}

func (o *PostUsersSignInDataUserPlexAccount) GetRestricted() *bool {
	if o == nil {
		return nil
	}
	return o.Restricted
}

func (o *PostUsersSignInDataUserPlexAccount) GetRoles() []string {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *PostUsersSignInDataUserPlexAccount) GetScrobbleTypes() string {
	if o == nil {
		return ""
	}
	return o.ScrobbleTypes
}

func (o *PostUsersSignInDataUserPlexAccount) GetServices() []PostUsersSignInDataServices {
	if o == nil {
		return []PostUsersSignInDataServices{}
	}
	return o.Services
}

func (o *PostUsersSignInDataUserPlexAccount) GetSubscription() PostUsersSignInDataSubscription {
	if o == nil {
		return PostUsersSignInDataSubscription{}
	}
	return o.Subscription
}

func (o *PostUsersSignInDataUserPlexAccount) GetSubscriptionDescription() *string {
	if o == nil {
		return nil
	}
	return o.SubscriptionDescription
}

func (o *PostUsersSignInDataUserPlexAccount) GetSubscriptions() []PostUsersSignInDataAuthenticationSubscription {
	if o == nil {
		return []PostUsersSignInDataAuthenticationSubscription{}
	}
	return o.Subscriptions
}

func (o *PostUsersSignInDataUserPlexAccount) GetThumb() string {
	if o == nil {
		return ""
	}
	return o.Thumb
}

func (o *PostUsersSignInDataUserPlexAccount) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *PostUsersSignInDataUserPlexAccount) GetTwoFactorEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.TwoFactorEnabled
}

func (o *PostUsersSignInDataUserPlexAccount) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *PostUsersSignInDataUserPlexAccount) GetUUID() string {
	if o == nil {
		return ""
	}
	return o.UUID
}

func (o *PostUsersSignInDataUserPlexAccount) GetAttributionPartner() *string {
	if o == nil {
		return nil
	}
	return o.AttributionPartner
}

func (o *PostUsersSignInDataUserPlexAccount) GetPastSubscriptions() []PastSubscription {
	if o == nil {
		return []PastSubscription{}
	}
	return o.PastSubscriptions
}

func (o *PostUsersSignInDataUserPlexAccount) GetTrials() []Trials {
	if o == nil {
		return []Trials{}
	}
	return o.Trials
}

type PostUsersSignInDataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns the user account data with a valid auth token
	UserPlexAccount *PostUsersSignInDataUserPlexAccount
}

func (o *PostUsersSignInDataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostUsersSignInDataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostUsersSignInDataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostUsersSignInDataResponse) GetUserPlexAccount() *PostUsersSignInDataUserPlexAccount {
	if o == nil {
		return nil
	}
	return o.UserPlexAccount
}
