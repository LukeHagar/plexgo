# github.com/LukeHagar/plexgo

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>

<!-- Start Summary [summary] -->
## Summary

Plex Media Server: OpenAPI specification for the Plex Media Server (PMS) API and the plex.tv cloud API.

## Base URLs

- **PMS (local server)**: `http(s)://{host}:{port}` — Most endpoints in this spec target the local PMS.
- **plex.tv v2**: `https://plex.tv/api/v2` — Authentication, account, and social endpoints.
- **plex.tv v1 (legacy)**: `https://plex.tv/api` — Legacy XML endpoints (friends, home users, claims).
- **Cloud providers**: `https://discover.provider.plex.tv`, `https://metadata.provider.plex.tv`, etc.

Endpoints that target plex.tv or cloud providers declare an override `servers` array.

## Authentication

- **X-Plex-Token**: Pass via the `X-Plex-Token` header on every request. It may also be passed as a query parameter (`?X-Plex-Token=...`) on all endpoints.
- **X-Plex-Client-Identifier**: Mandatory for OAuth PIN flow (`/pins`) and JWT device registration. Must be a unique, persistent identifier for the client application.
- **OAuth PIN Flow**: `POST /pins` → user visits `https://plex.tv/link` → `GET /pins/{pinId}` → obtain `authToken`.

## Response Formats

- **PMS endpoints**: Return XML by default. Send `Accept: application/json` to receive JSON.
- **plex.tv v2**: Returns JSON by default.
- **Legacy v1 endpoints** (`/pins.xml`, `/api/resources`, `/api/users/`): Return XML only.

## Rate Limiting

plex.tv auth endpoints (PIN creation, sign-in) enforce rate limits. Clients should implement exponential backoff and reuse tokens rather than re-authenticating on every request.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/LukeHagar/plexgo](#githubcomlukehagarplexgo)
  * [Base URLs](#base-urls)
  * [Authentication](#authentication)
  * [Response Formats](#response-formats)
  * [Rate Limiting](#rate-limiting)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Authentication](#authentication-1)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/LukeHagar/plexgo
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithAccepts(components.AcceptsApplicationXML),
		plexgo.WithClientIdentifier("abc123"),
		plexgo.WithProduct("Plex for Roku"),
		plexgo.WithVersion("2.4.1"),
		plexgo.WithPlatform("Roku"),
		plexgo.WithPlatformVersion("4.3 build 1057"),
		plexgo.WithDevice("Roku 3"),
		plexgo.WithModel("4200X"),
		plexgo.WithDeviceVendor("Roku"),
		plexgo.WithDeviceName("Living Room TV"),
		plexgo.WithMarketplace("googlePlay"),
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Transcoder.StartTranscodeSession(ctx, operations.StartTranscodeSessionRequest{
		TranscodeType:             components.TranscodeTypeMusic,
		AdvancedSubtitles:         components.AdvancedSubtitlesBurn.ToPointer(),
		Extension:                 operations.ExtensionMpd,
		AudioBoost:                plexgo.Pointer[int64](50),
		AudioChannelCount:         plexgo.Pointer[int64](5),
		AutoAdjustQuality:         components.BoolIntTrue.ToPointer(),
		AutoAdjustSubtitle:        components.BoolIntTrue.ToPointer(),
		DirectPlay:                components.BoolIntTrue.ToPointer(),
		DirectStream:              components.BoolIntTrue.ToPointer(),
		DirectStreamAudio:         components.BoolIntTrue.ToPointer(),
		DisableResolutionRotation: components.BoolIntTrue.ToPointer(),
		HasMDE:                    components.BoolIntTrue.ToPointer(),
		Location:                  operations.StartTranscodeSessionQueryParamLocationWan.ToPointer(),
		MediaBufferSize:           plexgo.Pointer[int64](102400),
		MediaIndex:                plexgo.Pointer[int64](0),
		MusicBitrate:              plexgo.Pointer[int64](5000),
		Offset:                    plexgo.Pointer[float64](90.5),
		PartIndex:                 plexgo.Pointer[int64](0),
		Path:                      plexgo.Pointer("/library/metadata/151671"),
		PeakBitrate:               plexgo.Pointer[int64](12000),
		PhotoResolution:           plexgo.Pointer("1080x1080"),
		Protocol:                  operations.StartTranscodeSessionQueryParamProtocolDash.ToPointer(),
		SecondsPerSegment:         plexgo.Pointer[int64](5),
		SubtitleSize:              plexgo.Pointer[int64](50),
		Subtitles:                 operations.StartTranscodeSessionQueryParamSubtitlesBurn.ToPointer(),
		VideoResolution:           plexgo.Pointer("1080x1080"),
		Copyts:                    components.BoolIntTrue.ToPointer(),
		VideoBitrate:              plexgo.Pointer[int64](12000),
		VideoQuality:              plexgo.Pointer[int64](50),
		XPlexClientProfileExtra:   plexgo.Pointer("add-limitation(scope=videoCodec&scopeName=*&type=upperBound&name=video.frameRate&value=60&replace=true)+append-transcode-target-codec(type=videoProfile&context=streaming&videoCodec=h264%2Chevc&audioCodec=aac&protocol=dash)"),
		XPlexClientProfileName:    plexgo.Pointer("generic"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.TwoHundredApplicationVndAppleMpegurlBinaryResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Activities](docs/sdks/activities/README.md)

* [ListActivities](docs/sdks/activities/README.md#listactivities) - Get all activities
* [CancelActivity](docs/sdks/activities/README.md#cancelactivity) - Cancel a running activity

### [Authentication](docs/sdks/authentication/README.md)

* [RegisterDeviceJWK](docs/sdks/authentication/README.md#registerdevicejwk) - Register Device JWK
* [GetAuthKeys](docs/sdks/authentication/README.md#getauthkeys) - Get Auth Keys
* [GetAuthNonce](docs/sdks/authentication/README.md#getauthnonce) - Get Auth Nonce
* [ExchangeJWTToken](docs/sdks/authentication/README.md#exchangejwttoken) - Exchange JWT Token
* [GetClaimToken](docs/sdks/authentication/README.md#getclaimtoken) - Get Claim Token
* [GetFeatures](docs/sdks/authentication/README.md#getfeatures) - Get Features
* [Ping](docs/sdks/authentication/README.md#ping) - Ping the server
* [CreateOAuthPin](docs/sdks/authentication/README.md#createoauthpin) - Create OAuth PIN
* [CreateLegacyPin](docs/sdks/authentication/README.md#createlegacypin) - Create Legacy PIN
* [LinkOAuthPin](docs/sdks/authentication/README.md#linkoauthpin) - Link OAuth PIN
* [GetServerAccessTokens](docs/sdks/authentication/README.md#getserveraccesstokens) - Get Server Access Tokens
* [GetTokenDetails](docs/sdks/authentication/README.md#gettokendetails) - Get Token Details
* [ChangePassword](docs/sdks/authentication/README.md#changepassword) - Change Password
* [PostUsersSignInData](docs/sdks/authentication/README.md#postuserssignindata) - Get User Sign In Data
* [SignOut](docs/sdks/authentication/README.md#signout) - Sign Out
* [SwitchHomeUser](docs/sdks/authentication/README.md#switchhomeuser) - Switch Home User
* [GetOAuthPin](docs/sdks/authentication/README.md#getoauthpin) - Get OAuth PIN Status

### [Butler](docs/sdks/butler/README.md)

* [StopTasks](docs/sdks/butler/README.md#stoptasks) - Stop all Butler tasks
* [GetTasks](docs/sdks/butler/README.md#gettasks) - Get all Butler tasks
* [StartTasks](docs/sdks/butler/README.md#starttasks) - Start all Butler tasks
* [StopTask](docs/sdks/butler/README.md#stoptask) - Stop a single Butler task
* [StartTask](docs/sdks/butler/README.md#starttask) - Start a single Butler task

### [Collections](docs/sdks/collections/README.md)

* [CreateCollection](docs/sdks/collections/README.md#createcollection) - Create collection

### [Content](docs/sdks/content/README.md)

* [GetCollectionItems](docs/sdks/content/README.md#getcollectionitems) - Get items in a collection
* [GetMetadataItem](docs/sdks/content/README.md#getmetadataitem) - Get a metadata item
* [GetAlbums](docs/sdks/content/README.md#getalbums) - Set section albums
* [ListContent](docs/sdks/content/README.md#listcontent) - Get items in the section
* [GetAllLeaves](docs/sdks/content/README.md#getallleaves) - Set section leaves
* [GetArts](docs/sdks/content/README.md#getarts) - Set section artwork
* [GetCategories](docs/sdks/content/README.md#getcategories) - Set section categories
* [GetCluster](docs/sdks/content/README.md#getcluster) - Set section clusters
* [GetSonicPath](docs/sdks/content/README.md#getsonicpath) - Similar tracks to transition from one to another
* [GetFolders](docs/sdks/content/README.md#getfolders) - Get all folder locations
* [ListMoments](docs/sdks/content/README.md#listmoments) - Set section moments
* [GetSonicallySimilar](docs/sdks/content/README.md#getsonicallysimilar) - The nearest audio tracks
* [GetCollectionImage](docs/sdks/content/README.md#getcollectionimage) - Get a collection's image

### [Devices](docs/sdks/devices/README.md)

* [GetAvailableGrabbers](docs/sdks/devices/README.md#getavailablegrabbers) - Get available grabbers
* [ListDevices](docs/sdks/devices/README.md#listdevices) - Get all devices
* [AddDevice](docs/sdks/devices/README.md#adddevice) - Add a device
* [DiscoverDevices](docs/sdks/devices/README.md#discoverdevices) - Tell grabbers to discover devices
* [RemoveDevice](docs/sdks/devices/README.md#removedevice) - Remove a device
* [GetDeviceDetails](docs/sdks/devices/README.md#getdevicedetails) - Get device details
* [ModifyDevice](docs/sdks/devices/README.md#modifydevice) - Enable or disable a device
* [SetChannelmap](docs/sdks/devices/README.md#setchannelmap) - Set a device's channel mapping
* [GetDevicesChannels](docs/sdks/devices/README.md#getdeviceschannels) - Get a device's channels
* [SetDevicePreferences](docs/sdks/devices/README.md#setdevicepreferences) - Set device preferences
* [StopScan](docs/sdks/devices/README.md#stopscan) - Tell a device to stop scanning for channels
* [Scan](docs/sdks/devices/README.md#scan) - Tell a device to scan for channels
* [GetThumb](docs/sdks/devices/README.md#getthumb) - Get device thumb

### [DownloadQueue](docs/sdks/downloadqueue/README.md)

* [CreateDownloadQueue](docs/sdks/downloadqueue/README.md#createdownloadqueue) - Create download queue
* [GetDownloadQueue](docs/sdks/downloadqueue/README.md#getdownloadqueue) - Get a download queue
* [AddDownloadQueueItems](docs/sdks/downloadqueue/README.md#adddownloadqueueitems) - Add to download queue
* [ListDownloadQueueItems](docs/sdks/downloadqueue/README.md#listdownloadqueueitems) - Get download queue items
* [GetItemDecision](docs/sdks/downloadqueue/README.md#getitemdecision) - Grab download queue item decision
* [GetDownloadQueueMedia](docs/sdks/downloadqueue/README.md#getdownloadqueuemedia) - Grab download queue media
* [RemoveDownloadQueueItems](docs/sdks/downloadqueue/README.md#removedownloadqueueitems) - Delete download queue items
* [GetDownloadQueueItems](docs/sdks/downloadqueue/README.md#getdownloadqueueitems) - Get download queue items
* [RestartProcessingDownloadQueueItems](docs/sdks/downloadqueue/README.md#restartprocessingdownloadqueueitems) - Restart processing of items from the decision

### [DVRs](docs/sdks/dvrs/README.md)

* [ListDVRs](docs/sdks/dvrs/README.md#listdvrs) - Get DVRs
* [CreateDVR](docs/sdks/dvrs/README.md#createdvr) - Create a DVR
* [DeleteDVR](docs/sdks/dvrs/README.md#deletedvr) - Delete a single DVR
* [GetDVR](docs/sdks/dvrs/README.md#getdvr) - Get a single DVR
* [PatchDVRSettings](docs/sdks/dvrs/README.md#patchdvrsettings) - Update DVR Settings
* [UpdateDVRSettings](docs/sdks/dvrs/README.md#updatedvrsettings) - Update DVR Settings
* [GetDVRChannels](docs/sdks/dvrs/README.md#getdvrchannels) - Get DVR Channels
* [GetDVRGuide](docs/sdks/dvrs/README.md#getdvrguide) - Get DVR Guide
* [DeleteLineup](docs/sdks/dvrs/README.md#deletelineup) - Delete a DVR Lineup
* [AddLineup](docs/sdks/dvrs/README.md#addlineup) - Add a DVR Lineup
* [SetDVRPreferences](docs/sdks/dvrs/README.md#setdvrpreferences) - Set DVR preferences
* [StopDVRReload](docs/sdks/dvrs/README.md#stopdvrreload) - Tell a DVR to stop reloading program guide
* [ReloadGuide](docs/sdks/dvrs/README.md#reloadguide) - Tell a DVR to reload program guide
* [TuneChannel](docs/sdks/dvrs/README.md#tunechannel) - Tune a channel on a DVR
* [RemoveDeviceFromDVR](docs/sdks/dvrs/README.md#removedevicefromdvr) - Remove a device from an existing DVR
* [AddDeviceToDVR](docs/sdks/dvrs/README.md#adddevicetodvr) - Add a device to an existing DVR

### [Epg](docs/sdks/epg/README.md)

* [ComputeChannelMap](docs/sdks/epg/README.md#computechannelmap) - Compute the best channel map
* [GetChannels](docs/sdks/epg/README.md#getchannels) - Get channels for a lineup
* [GetCountries](docs/sdks/epg/README.md#getcountries) - Get all countries
* [GetEPGGuide](docs/sdks/epg/README.md#getepgguide) - Get EPG Guide
* [GetAllLanguages](docs/sdks/epg/README.md#getalllanguages) - Get all languages
* [GetLineup](docs/sdks/epg/README.md#getlineup) - Compute the best lineup
* [GetLineupChannels](docs/sdks/epg/README.md#getlineupchannels) - Get the channels for multiple lineups
* [SearchEPG](docs/sdks/epg/README.md#searchepg) - Search EPG
* [GetCountriesLineups](docs/sdks/epg/README.md#getcountrieslineups) - Get lineups for a country via postal code
* [GetCountryRegions](docs/sdks/epg/README.md#getcountryregions) - Get regions for a country
* [ListLineups](docs/sdks/epg/README.md#listlineups) - Get lineups for a region

### [Events](docs/sdks/events/README.md)

* [GetNotifications](docs/sdks/events/README.md#getnotifications) - Connect to Eventsource
* [ConnectWebSocket](docs/sdks/events/README.md#connectwebsocket) - Connect to WebSocket
* [GetWebsocketNotifications](docs/sdks/events/README.md#getwebsocketnotifications) - Get WebSocket Notifications

### [General](docs/sdks/general/README.md)

* [GetServerInfo](docs/sdks/general/README.md#getserverinfo) - Get PMS info
* [GetSystemAccounts](docs/sdks/general/README.md#getsystemaccounts) - Get System Accounts
* [GetUserWebhooks](docs/sdks/general/README.md#getuserwebhooks) - User Webhooks
* [AddUserWebhook](docs/sdks/general/README.md#adduserwebhook) - Add User Webhook
* [GetClients](docs/sdks/general/README.md#getclients) - Get Clients
* [GetCloudServer](docs/sdks/general/README.md#getcloudserver) - Get Cloud Server
* [GetSystemDevices](docs/sdks/general/README.md#getsystemdevices) - Get System Devices
* [GetDiagnostics](docs/sdks/general/README.md#getdiagnostics) - Get Diagnostics
* [DownloadDatabaseDiagnostics](docs/sdks/general/README.md#downloaddatabasediagnostics) - Download Database Diagnostics
* [DownloadLogBundle](docs/sdks/general/README.md#downloadlogbundle) - Download Log Bundle
* [GetGeoIP](docs/sdks/general/README.md#getgeoip) - Get GeoIP
* [GetIdentity](docs/sdks/general/README.md#getidentity) - Get PMS identity
* [GetIP](docs/sdks/general/README.md#getip) - Get IP
* [ClaimServer](docs/sdks/general/README.md#claimserver) - Claim Server
* [RefreshReachability](docs/sdks/general/README.md#refreshreachability) - Refresh Reachability
* [GetSourceConnectionInformation](docs/sdks/general/README.md#getsourceconnectioninformation) - Get Source Connection Information
* [CreateTransientToken](docs/sdks/general/README.md#createtransienttoken) - Get Transient Tokens
* [GetLocalServers](docs/sdks/general/README.md#getlocalservers) - Get Local Servers
* [BrowseFilesystem](docs/sdks/general/README.md#browsefilesystem) - Browse Filesystem
* [GetBandwidthStatistics](docs/sdks/general/README.md#getbandwidthstatistics) - Get Bandwidth Statistics
* [GetResourceStatistics](docs/sdks/general/README.md#getresourcestatistics) - Get Resource Statistics
* [GetSyncStatus](docs/sdks/general/README.md#getsyncstatus) - Get Sync Status
* [GetSyncItems](docs/sdks/general/README.md#getsyncitems) - Get Sync Items
* [GetSyncQueue](docs/sdks/general/README.md#getsyncqueue) - Get Sync Queue
* [RefreshSyncContent](docs/sdks/general/README.md#refreshsynccontent) - Refresh Sync Content
* [RefreshSyncLists](docs/sdks/general/README.md#refreshsynclists) - Refresh Sync Lists
* [GetSyncTranscodeQueue](docs/sdks/general/README.md#getsynctranscodequeue) - Get Sync Transcode Queue
* [GetMetadataAgents](docs/sdks/general/README.md#getmetadataagents) - Get Metadata Agents
* [GetSystemSettings](docs/sdks/general/README.md#getsystemsettings) - Get System Settings
* [CheckForSystemUpdates](docs/sdks/general/README.md#checkforsystemupdates) - Check for System Updates
* [GetWebhooks](docs/sdks/general/README.md#getwebhooks) - Get Webhooks
* [AddWebhook](docs/sdks/general/README.md#addwebhook) - Add Webhook
* [GetPlexDownloads](docs/sdks/general/README.md#getplexdownloads) - Get Plex Downloads
* [BrowseFilesystemPath](docs/sdks/general/README.md#browsefilesystempath) - Browse Filesystem Path
* [GetSyncItem](docs/sdks/general/README.md#getsyncitem) - Get Sync Item
* [GetMetadataAgentDetails](docs/sdks/general/README.md#getmetadataagentdetails) - Get Metadata Agent Details

### [Hubs](docs/sdks/hubs/README.md)

* [GetAllHubs](docs/sdks/hubs/README.md#getallhubs) - Get global hubs
* [GetContinueWatching](docs/sdks/hubs/README.md#getcontinuewatching) - Get the continue watching hub
* [GetContinueWatchingItems](docs/sdks/hubs/README.md#getcontinuewatchingitems) - Get Continue Watching Items
* [GetHomeRecentlyAdded](docs/sdks/hubs/README.md#gethomerecentlyadded) - Get home hubs Recently Added
* [GetHubItems](docs/sdks/hubs/README.md#gethubitems) - Get a hub's items
* [GetPromotedHubs](docs/sdks/hubs/README.md#getpromotedhubs) - Get the hubs which are promoted
* [GetMetadataHubs](docs/sdks/hubs/README.md#getmetadatahubs) - Get hubs for section by metadata item
* [GetPostplayHubs](docs/sdks/hubs/README.md#getpostplayhubs) - Get postplay hubs
* [GetRelatedHubs](docs/sdks/hubs/README.md#getrelatedhubs) - Get related hubs
* [GetSectionHubs](docs/sdks/hubs/README.md#getsectionhubs) - Get section hubs
* [ResetSectionDefaults](docs/sdks/hubs/README.md#resetsectiondefaults) - Reset hubs to defaults
* [ListHubs](docs/sdks/hubs/README.md#listhubs) - Get hubs
* [CreateCustomHub](docs/sdks/hubs/README.md#createcustomhub) - Create a custom hub
* [MoveHub](docs/sdks/hubs/README.md#movehub) - Move Hub
* [DeleteCustomHub](docs/sdks/hubs/README.md#deletecustomhub) - Delete a custom hub
* [UpdateHubVisibility](docs/sdks/hubs/README.md#updatehubvisibility) - Change hub visibility

### [Library](docs/sdks/library/README.md)

* [GetRootLibrary](docs/sdks/library/README.md#getrootlibrary) - Get Root Library
* [GetLibraryItems](docs/sdks/library/README.md#getlibraryitems) - Get all items in library
* [DeleteCaches](docs/sdks/library/README.md#deletecaches) - Delete library caches
* [CleanBundles](docs/sdks/library/README.md#cleanbundles) - Clean bundles
* [IngestTransientItem](docs/sdks/library/README.md#ingesttransientitem) - Ingest a transient item
* [GetLibraryMatches](docs/sdks/library/README.md#getlibrarymatches) - Get library matches
* [OptimizeLibrary](docs/sdks/library/README.md#optimizelibrary) - Get Optimize Library
* [OptimizeLibraryPost](docs/sdks/library/README.md#optimizelibrarypost) - Optimize Library
* [OptimizeDatabase](docs/sdks/library/README.md#optimizedatabase) - Optimize the Database
* [GetRandomArtwork](docs/sdks/library/README.md#getrandomartwork) - Get random artwork
* [GetRecentlyAddedGlobal](docs/sdks/library/README.md#getrecentlyaddedglobal) - Get Global Recently Added
* [GetLibrarySectionsFallback](docs/sdks/library/README.md#getlibrarysectionsfallback) - Get Library Sections (Fallback)
* [GetSections](docs/sdks/library/README.md#getsections) - Get library sections (main Media Provider Only)
* [AddSection](docs/sdks/library/README.md#addsection) - Add a library section
* [StopAllRefreshes](docs/sdks/library/README.md#stopallrefreshes) - Stop refresh
* [GetSectionsPrefs](docs/sdks/library/README.md#getsectionsprefs) - Get section prefs
* [RefreshSectionsMetadata](docs/sdks/library/README.md#refreshsectionsmetadata) - Refresh all sections
* [GetTags](docs/sdks/library/README.md#gettags) - Get all library tags of a type
* [UploadArt](docs/sdks/library/README.md#uploadart) - Upload media art Art
* [GetMetadataChildren](docs/sdks/library/README.md#getmetadatachildren) - Get Metadata Children
* [ComputeSonicPath](docs/sdks/library/README.md#computesonicpath) - Compute Sonic Path
* [GetMetadataGrandchildren](docs/sdks/library/README.md#getmetadatagrandchildren) - Get Metadata Grandchildren
* [GetMetadataGrandparent](docs/sdks/library/README.md#getmetadatagrandparent) - Get Metadata Grandparent
* [GetNearestMetadata](docs/sdks/library/README.md#getnearestmetadata) - Get Nearest Metadata
* [GetMetadataOnDeck](docs/sdks/library/README.md#getmetadataondeck) - Get Metadata On Deck
* [GetMetadataParent](docs/sdks/library/README.md#getmetadataparent) - Get Metadata Parent
* [UploadPoster](docs/sdks/library/README.md#uploadposter) - Upload media art Poster
* [GetMetadataReviews](docs/sdks/library/README.md#getmetadatareviews) - Get Metadata Reviews
* [DeleteMetadataItem](docs/sdks/library/README.md#deletemetadataitem) - Delete a metadata item
* [EditMetadataItem](docs/sdks/library/README.md#editmetadataitem) - Edit a metadata item
* [DetectAds](docs/sdks/library/README.md#detectads) - Ad-detect an item
* [GetAllItemLeaves](docs/sdks/library/README.md#getallitemleaves) - Get the leaves of an item
* [AnalyzeMetadata](docs/sdks/library/README.md#analyzemetadata) - Analyze an item
* [GenerateThumbs](docs/sdks/library/README.md#generatethumbs) - Generate thumbs of chapters for an item
* [DetectCredits](docs/sdks/library/README.md#detectcredits) - Credit detect a metadata item
* [GetExtras](docs/sdks/library/README.md#getextras) - Get an item's extras
* [AddExtras](docs/sdks/library/README.md#addextras) - Add to an item's extras
* [GetFile](docs/sdks/library/README.md#getfile) - Get a file from a metadata or media bundle
* [StartBifGeneration](docs/sdks/library/README.md#startbifgeneration) - Start BIF generation of an item
* [DetectIntros](docs/sdks/library/README.md#detectintros) - Intro detect an item
* [CreateMarker](docs/sdks/library/README.md#createmarker) - Create a marker
* [MatchItem](docs/sdks/library/README.md#matchitem) - Match a metadata item
* [ListMatches](docs/sdks/library/README.md#listmatches) - Get metadata matches for an item
* [MergeItems](docs/sdks/library/README.md#mergeitems) - Merge a metadata item
* [SetItemPreferences](docs/sdks/library/README.md#setitempreferences) - Set metadata preferences
* [RefreshItemsMetadata](docs/sdks/library/README.md#refreshitemsmetadata) - Refresh a metadata item
* [GetRelatedItems](docs/sdks/library/README.md#getrelateditems) - Get related items
* [ListSimilar](docs/sdks/library/README.md#listsimilar) - Get similar items
* [SplitItem](docs/sdks/library/README.md#splititem) - Split a metadata item
* [GetSubtitles](docs/sdks/library/README.md#getsubtitles) - Get subtitles
* [GetItemTree](docs/sdks/library/README.md#getitemtree) - Get metadata items as a tree
* [Unmatch](docs/sdks/library/README.md#unmatch) - Unmatch a metadata item
* [ListTopUsers](docs/sdks/library/README.md#listtopusers) - Get metadata top users
* [DetectVoiceActivity](docs/sdks/library/README.md#detectvoiceactivity) - Detect voice activity
* [GetAugmentationStatus](docs/sdks/library/README.md#getaugmentationstatus) - Get augmentation status
* [SetStreamSelection](docs/sdks/library/README.md#setstreamselection) - Set stream selection
* [GetPerson](docs/sdks/library/README.md#getperson) - Get person details
* [ListPersonMedia](docs/sdks/library/README.md#listpersonmedia) - Get media for a person
* [DeleteLibrarySection](docs/sdks/library/README.md#deletelibrarysection) - Delete a library section
* [GetLibraryDetails](docs/sdks/library/README.md#getlibrarydetails) - Get a library section by id
* [EditSection](docs/sdks/library/README.md#editsection) - Edit a library section
* [GetSectionAgents](docs/sdks/library/README.md#getsectionagents) - Get Section Agents
* [UpdateItems](docs/sdks/library/README.md#updateitems) - Set the fields of the filtered items
* [StartAnalysis](docs/sdks/library/README.md#startanalysis) - Analyze a section
* [GetSectionArtists](docs/sdks/library/README.md#getsectionartists) - Get Section Artists
* [Autocomplete](docs/sdks/library/README.md#autocomplete) - Get autocompletions for search
* [GetByContentRating](docs/sdks/library/README.md#getbycontentrating) - Get By Content Rating
* [GetByDecade](docs/sdks/library/README.md#getbydecade) - Get By Decade
* [GetByFolder](docs/sdks/library/README.md#getbyfolder) - Get By Folder
* [GetByResolution](docs/sdks/library/README.md#getbyresolution) - Get By Resolution
* [GetByYear](docs/sdks/library/README.md#getbyyear) - Get By Year
* [GetSectionClips](docs/sdks/library/README.md#getsectionclips) - Get Section Clips
* [GetCollections](docs/sdks/library/README.md#getcollections) - Get collections in a section
* [GetCommon](docs/sdks/library/README.md#getcommon) - Get common fields for items
* [GetSectionEdit](docs/sdks/library/README.md#getsectionedit) - Edit Section
* [EditLibrarySection](docs/sdks/library/README.md#editlibrarysection) - Edit Section
* [EmptyTrash](docs/sdks/library/README.md#emptytrash) - Get Empty Trash
* [EmptyTrashPost](docs/sdks/library/README.md#emptytrashpost) - Empty Trash
* [EmptyTrashPut](docs/sdks/library/README.md#emptytrashput) - Empty section trash
* [GetSectionEpisodes](docs/sdks/library/README.md#getsectionepisodes) - Get Section Episodes
* [GetSectionFilters](docs/sdks/library/README.md#getsectionfilters) - Get section filters
* [GetFirstCharacters](docs/sdks/library/README.md#getfirstcharacters) - Get list of first characters
* [GetLibrarySectionHubs](docs/sdks/library/README.md#getlibrarysectionhubs) - Get Section Hubs
* [DeleteIndexes](docs/sdks/library/README.md#deleteindexes) - Delete section indexes
* [DeleteIntros](docs/sdks/library/README.md#deleteintros) - Delete section intro markers
* [GetSectionLabels](docs/sdks/library/README.md#getsectionlabels) - Get Section Labels
* [MatchSectionItems](docs/sdks/library/README.md#matchsectionitems) - Match Section Items
* [MoveSection](docs/sdks/library/README.md#movesection) - Move Section
* [GetSectionMovies](docs/sdks/library/README.md#getsectionmovies) - Get Section Movies
* [GetNewestForSection](docs/sdks/library/README.md#getnewestforsection) - Get Newest for Section
* [GetOnDeckForSection](docs/sdks/library/README.md#getondeckforsection) - Get On Deck for Section
* [OptimizeSection](docs/sdks/library/README.md#optimizesection) - Get Optimize Section
* [OptimizeSectionPost](docs/sdks/library/README.md#optimizesectionpost) - Optimize Section
* [GetSectionPhotos](docs/sdks/library/README.md#getsectionphotos) - Get Section Photos
* [GetSectionPlaylists](docs/sdks/library/README.md#getsectionplaylists) - Get Section Playlists
* [GetSectionPreferences](docs/sdks/library/README.md#getsectionpreferences) - Get section prefs
* [SetSectionPreferences](docs/sdks/library/README.md#setsectionpreferences) - Set section prefs
* [GetRecentlyAddedForSection](docs/sdks/library/README.md#getrecentlyaddedforsection) - Get Recently Added for Section
* [CancelRefresh](docs/sdks/library/README.md#cancelrefresh) - Cancel section refresh
* [RefreshSection](docs/sdks/library/README.md#refreshsection) - Get Refresh Section
* [RefreshSectionPost](docs/sdks/library/README.md#refreshsectionpost) - Refresh Section
* [SearchSection](docs/sdks/library/README.md#searchsection) - Search Section
* [GetSectionSettings](docs/sdks/library/README.md#getsectionsettings) - Get Section Settings
* [GetSectionShows](docs/sdks/library/README.md#getsectionshows) - Get Section Shows
* [GetAvailableSorts](docs/sdks/library/README.md#getavailablesorts) - Get a section sorts
* [GetSectionTags](docs/sdks/library/README.md#getsectiontags) - Get Section Tags
* [GetSectionTimeline](docs/sdks/library/README.md#getsectiontimeline) - Get Section Timeline
* [UnmatchSectionItems](docs/sdks/library/README.md#unmatchsectionitems) - Unmatch Section Items
* [GetUnwatchedForSection](docs/sdks/library/README.md#getunwatchedforsection) - Get Unwatched for Section
* [GetStreamLevels](docs/sdks/library/README.md#getstreamlevels) - Get loudness about a stream in json
* [GetStreamLoudness](docs/sdks/library/README.md#getstreamloudness) - Get loudness about a stream
* [GetChapterImage](docs/sdks/library/README.md#getchapterimage) - Get a chapter image
* [SetItemArtwork](docs/sdks/library/README.md#setitemartwork) - Set an item's artwork, theme, etc
* [UpdateItemArtwork](docs/sdks/library/README.md#updateitemartwork) - Set an item's artwork, theme, etc
* [DeleteMarker](docs/sdks/library/README.md#deletemarker) - Delete a marker
* [EditMarker](docs/sdks/library/README.md#editmarker) - Edit a marker
* [DeleteMediaItem](docs/sdks/library/README.md#deletemediaitem) - Delete a media item
* [GetPartIndex](docs/sdks/library/README.md#getpartindex) - Get BIF index for a part
* [DeleteCollection](docs/sdks/library/README.md#deletecollection) - Delete a collection
* [GetSectionImage](docs/sdks/library/README.md#getsectionimage) - Get a section composite image
* [DeleteStream](docs/sdks/library/README.md#deletestream) - Delete a stream
* [GetStream](docs/sdks/library/README.md#getstream) - Get a stream
* [SetStreamOffset](docs/sdks/library/README.md#setstreamoffset) - Set a stream offset
* [GetItemArtwork](docs/sdks/library/README.md#getitemartwork) - Get an item's artwork, theme, etc
* [GetMediaPart](docs/sdks/library/README.md#getmediapart) - Get a media part
* [GetImageFromBif](docs/sdks/library/README.md#getimagefrombif) - Get an image from part BIF

### [LibraryCollections](docs/sdks/librarycollections/README.md)

* [AddCollectionItems](docs/sdks/librarycollections/README.md#addcollectionitems) - Add items to a collection
* [UpdateCollectionItem](docs/sdks/librarycollections/README.md#updatecollectionitem) - Update an item in a collection
* [MoveCollectionItem](docs/sdks/librarycollections/README.md#movecollectionitem) - Reorder an item in the collection

### [LibraryPlaylists](docs/sdks/libraryplaylists/README.md)

* [CreatePlaylist](docs/sdks/libraryplaylists/README.md#createplaylist) - Create a Playlist
* [UploadPlaylist](docs/sdks/libraryplaylists/README.md#uploadplaylist) - Upload media art
* [DeletePlaylist](docs/sdks/libraryplaylists/README.md#deleteplaylist) - Delete a Playlist
* [UpdatePlaylist](docs/sdks/libraryplaylists/README.md#updateplaylist) - Editing a Playlist
* [GetPlaylistGenerators](docs/sdks/libraryplaylists/README.md#getplaylistgenerators) - Get a playlist's generators
* [ClearPlaylistItems](docs/sdks/libraryplaylists/README.md#clearplaylistitems) - Clearing a playlist
* [AddPlaylistItems](docs/sdks/libraryplaylists/README.md#addplaylistitems) - Adding to  a Playlist
* [DeletePlaylistItem](docs/sdks/libraryplaylists/README.md#deleteplaylistitem) - Delete a Generator
* [GetPlaylistGenerator](docs/sdks/libraryplaylists/README.md#getplaylistgenerator) - Get a playlist generator
* [ModifyPlaylistGenerator](docs/sdks/libraryplaylists/README.md#modifyplaylistgenerator) - Modify a Generator
* [GetPlaylistGeneratorItems](docs/sdks/libraryplaylists/README.md#getplaylistgeneratoritems) - Get a playlist generator's items
* [MovePlaylistItem](docs/sdks/libraryplaylists/README.md#moveplaylistitem) - Moving items in a playlist
* [RefreshPlaylist](docs/sdks/libraryplaylists/README.md#refreshplaylist) - Reprocess a generator

### [LiveTV](docs/sdks/livetv/README.md)

* [GetDVRRecordings](docs/sdks/livetv/README.md#getdvrrecordings) - Get DVR Recordings
* [GetSessions](docs/sdks/livetv/README.md#getsessions) - Get all sessions
* [GetDVRRecordingsByDVR](docs/sdks/livetv/README.md#getdvrrecordingsbydvr) - Get DVR Recordings by DVR
* [DeleteLiveTVSession](docs/sdks/livetv/README.md#deletelivetvsession) - Delete Live TV Session
* [GetLiveTVSession](docs/sdks/livetv/README.md#getlivetvsession) - Get a single session
* [GetSessionPlaylistIndex](docs/sdks/livetv/README.md#getsessionplaylistindex) - Get a session playlist index
* [GetSessionSegment](docs/sdks/livetv/README.md#getsessionsegment) - Get a single session segment

### [Log](docs/sdks/log/README.md)

* [WriteLog](docs/sdks/log/README.md#writelog) - Logging a multi-line message to the Plex Media Server log
* [WriteMessage](docs/sdks/log/README.md#writemessage) - Logging a single-line message to the Plex Media Server log
* [EnablePapertrail](docs/sdks/log/README.md#enablepapertrail) - Enabling Papertrail

### [PlayQueue](docs/sdks/playqueue/README.md)

* [CreatePlayQueue](docs/sdks/playqueue/README.md#createplayqueue) - Create a play queue
* [GetPlayQueue](docs/sdks/playqueue/README.md#getplayqueue) - Retrieve a play queue
* [AddToPlayQueue](docs/sdks/playqueue/README.md#addtoplayqueue) - Add a generator or playlist to a play queue
* [ClearPlayQueue](docs/sdks/playqueue/README.md#clearplayqueue) - Clear a play queue
* [ResetPlayQueue](docs/sdks/playqueue/README.md#resetplayqueue) - Reset a play queue
* [Shuffle](docs/sdks/playqueue/README.md#shuffle) - Shuffle a play queue
* [Unshuffle](docs/sdks/playqueue/README.md#unshuffle) - Unshuffle a play queue
* [DeletePlayQueueItem](docs/sdks/playqueue/README.md#deleteplayqueueitem) - Delete an item from a play queue
* [MovePlayQueueItem](docs/sdks/playqueue/README.md#moveplayqueueitem) - Move an item in a play queue

### [Playback](docs/sdks/playback/README.md)

* [GetProgress](docs/sdks/playback/README.md#getprogress) - Get Progress
* [RemoveFromContinueWatching](docs/sdks/playback/README.md#removefromcontinuewatching) - Remove From Continue Watching
* [PlayerAudioStream](docs/sdks/playback/README.md#playeraudiostream) - Player Audio Stream
* [PlayerMute](docs/sdks/playback/README.md#playermute) - Player Mute
* [PlayerPause](docs/sdks/playback/README.md#playerpause) - Player Pause
* [PlayerPlay](docs/sdks/playback/README.md#playerplay) - Player Play
* [PlayerPlayMedia](docs/sdks/playback/README.md#playerplaymedia) - Player Play Media
* [PlayerRefreshplayqueue](docs/sdks/playback/README.md#playerrefreshplayqueue) - Player Refresh Play Queue
* [PlayerSeek](docs/sdks/playback/README.md#playerseek) - Player Seek
* [PlayerSetParameters](docs/sdks/playback/README.md#playersetparameters) - Player Set Parameters
* [PlayerSetRating](docs/sdks/playback/README.md#playersetrating) - Player Set Rating
* [PlayerSetState](docs/sdks/playback/README.md#playersetstate) - Player Set State
* [PlayerSetStreams](docs/sdks/playback/README.md#playersetstreams) - Player Set Streams
* [PlayerSetTextStream](docs/sdks/playback/README.md#playersettextstream) - Player Set Text Stream
* [PlayerSetViewOffset](docs/sdks/playback/README.md#playersetviewoffset) - Player Set View Offset
* [PlayerSkipBy](docs/sdks/playback/README.md#playerskipby) - Player Skip By
* [PlayerSkipTo](docs/sdks/playback/README.md#playerskipto) - Player Skip To
* [PlayerStepback](docs/sdks/playback/README.md#playerstepback) - Player Step Back
* [PlayerStepforward](docs/sdks/playback/README.md#playerstepforward) - Player Step Forward
* [PlayerStop](docs/sdks/playback/README.md#playerstop) - Player Stop
* [PlayerSubtitleStream](docs/sdks/playback/README.md#playersubtitlestream) - Player Subtitle Stream
* [PlayerUnmute](docs/sdks/playback/README.md#playerunmute) - Player Unmute
* [PlayerVideoStream](docs/sdks/playback/README.md#playervideostream) - Player Video Stream
* [PlayerVolume](docs/sdks/playback/README.md#playervolume) - Player Volume
* [GetClientResources](docs/sdks/playback/README.md#getclientresources) - Get Client Resources
* [PlayerPollTimeline](docs/sdks/playback/README.md#playerpolltimeline) - Player Poll Timeline

### [Playlist](docs/sdks/playlist/README.md)

* [ListPlaylists](docs/sdks/playlist/README.md#listplaylists) - List playlists
* [GetPlaylist](docs/sdks/playlist/README.md#getplaylist) - Retrieve Playlist
* [GetPlaylistItems](docs/sdks/playlist/README.md#getplaylistitems) - Retrieve Playlist Contents

### [Playlists](docs/sdks/playlists/README.md)

* [DeletePlaylistByRatingKey](docs/sdks/playlists/README.md#deleteplaylistbyratingkey) - Delete Playlist

### [Plex](docs/sdks/plex/README.md)

* [GetServerResources](docs/sdks/plex/README.md#getserverresources) - Get Server Resources

### [Preferences](docs/sdks/preferences/README.md)

* [GetAllPreferences](docs/sdks/preferences/README.md#getallpreferences) - Get all preferences
* [SetPreferences](docs/sdks/preferences/README.md#setpreferences) - Set preferences
* [GetPreference](docs/sdks/preferences/README.md#getpreference) - Get a preferences

### [Provider](docs/sdks/provider/README.md)

* [AddToWatchlist](docs/sdks/provider/README.md#addtowatchlist) - Add to Watchlist
* [RemoveFromWatchlist](docs/sdks/provider/README.md#removefromwatchlist) - Remove from Watchlist
* [SearchDiscover](docs/sdks/provider/README.md#searchdiscover) - Search Discover
* [GetWatchlist](docs/sdks/provider/README.md#getwatchlist) - Get Watchlist
* [ListProviders](docs/sdks/provider/README.md#listproviders) - Get the list of available media providers
* [AddProvider](docs/sdks/provider/README.md#addprovider) - Add a media provider
* [RefreshProviders](docs/sdks/provider/README.md#refreshproviders) - Refresh media providers
* [DeleteMediaProvider](docs/sdks/provider/README.md#deletemediaprovider) - Delete a media provider

### [Rate](docs/sdks/rate/README.md)

* [SetRating](docs/sdks/rate/README.md#setrating) - Rate an item

### [Search](docs/sdks/search/README.md)

* [SearchHubs](docs/sdks/search/README.md#searchhubs) - Search Hub
* [VoiceSearchHubs](docs/sdks/search/README.md#voicesearchhubs) - Voice Search Hub

### [Status](docs/sdks/status/README.md)

* [ListSessions](docs/sdks/status/README.md#listsessions) - List Sessions
* [GetBackgroundTasks](docs/sdks/status/README.md#getbackgroundtasks) - Get background tasks
* [ListPlaybackHistory](docs/sdks/status/README.md#listplaybackhistory) - List Playback History
* [TerminateSession](docs/sdks/status/README.md#terminatesession) - Terminate a session
* [DeleteHistory](docs/sdks/status/README.md#deletehistory) - Delete Single History Item
* [GetHistoryItem](docs/sdks/status/README.md#gethistoryitem) - Get Single History Item

### [Subscriptions](docs/sdks/subscriptions/README.md)

* [GetAllSubscriptions](docs/sdks/subscriptions/README.md#getallsubscriptions) - Get all subscriptions
* [CreateSubscription](docs/sdks/subscriptions/README.md#createsubscription) - Create a subscription
* [ProcessSubscriptions](docs/sdks/subscriptions/README.md#processsubscriptions) - Process all subscriptions
* [GetScheduledRecordings](docs/sdks/subscriptions/README.md#getscheduledrecordings) - Get all scheduled recordings
* [GetTemplate](docs/sdks/subscriptions/README.md#gettemplate) - Get the subscription template
* [CancelGrab](docs/sdks/subscriptions/README.md#cancelgrab) - Cancel an existing grab
* [DeleteSubscription](docs/sdks/subscriptions/README.md#deletesubscription) - Delete a subscription
* [GetSubscription](docs/sdks/subscriptions/README.md#getsubscription) - Get a single subscription
* [EditSubscriptionPreferences](docs/sdks/subscriptions/README.md#editsubscriptionpreferences) - Edit a subscription
* [ReorderSubscription](docs/sdks/subscriptions/README.md#reordersubscription) - Re-order a subscription

### [Timeline](docs/sdks/timeline/README.md)

* [MarkPlayed](docs/sdks/timeline/README.md#markplayed) - Mark an item as played
* [Report](docs/sdks/timeline/README.md#report) - Report media timeline
* [Unscrobble](docs/sdks/timeline/README.md#unscrobble) - Mark an item as unplayed
* [GetConversionQueue](docs/sdks/timeline/README.md#getconversionqueue) - Get Conversion Queue

### [Transcoder](docs/sdks/transcoder/README.md)

* [TranscodeMusic](docs/sdks/transcoder/README.md#transcodemusic) - Transcode Music
* [TranscodeImage](docs/sdks/transcoder/README.md#transcodeimage) - Transcode an image
* [GetTranscodeSessions](docs/sdks/transcoder/README.md#gettranscodesessions) - Get Transcode Sessions
* [MakeDecision](docs/sdks/transcoder/README.md#makedecision) - Make a decision on media playback
* [TriggerFallback](docs/sdks/transcoder/README.md#triggerfallback) - Manually trigger a transcoder fallback
* [TranscodeSubtitles](docs/sdks/transcoder/README.md#transcodesubtitles) - Transcode subtitles
* [StartTranscodeSession](docs/sdks/transcoder/README.md#starttranscodesession) - Start A Transcoding Session
* [GetDASHSegment](docs/sdks/transcoder/README.md#getdashsegment) - Get DASH Segment
* [GetHLSSegment](docs/sdks/transcoder/README.md#gethlssegment) - Get HLS Segment

### [UltraBlur](docs/sdks/ultrablur/README.md)

* [GetColors](docs/sdks/ultrablur/README.md#getcolors) - Get UltraBlur Colors
* [GetImage](docs/sdks/ultrablur/README.md#getimage) - Get UltraBlur Image

### [Updater](docs/sdks/updater/README.md)

* [ApplyUpdates](docs/sdks/updater/README.md#applyupdates) - Applying updates
* [CheckUpdates](docs/sdks/updater/README.md#checkupdates) - Checking for updates
* [GetUpdatesStatus](docs/sdks/updater/README.md#getupdatesstatus) - Querying status of updates

### [Users](docs/sdks/users/README.md)

* [GetLegacyResources](docs/sdks/users/README.md#getlegacyresources) - Get Legacy Resources
* [GetLegacyUsers](docs/sdks/users/README.md#getlegacyusers) - Get Legacy Users
* [GetFriends](docs/sdks/users/README.md#getfriends) - Get Friends
* [GetHome](docs/sdks/users/README.md#gethome) - Get home hubs
* [GetHomeUsers](docs/sdks/users/README.md#gethomeusers) - Get home hubs Users
* [CreateHomeUser](docs/sdks/users/README.md#createhomeuser) - Create Home User
* [GetMyPlexAccount](docs/sdks/users/README.md#getmyplexaccount) - Get MyPlex Account
* [GetUserServer](docs/sdks/users/README.md#getuserserver) - Get User Server Association
* [GetServerUserFeatures](docs/sdks/users/README.md#getserveruserfeatures) - Get Server User Features
* [ShareServer](docs/sdks/users/README.md#shareserver) - Share Server
* [UpdateViewStateSync](docs/sdks/users/README.md#updateviewstatesync) - Update View State Sync
* [GetUsers](docs/sdks/users/README.md#getusers) - Get list of all connected users
* [GetAccountXML](docs/sdks/users/README.md#getaccountxml) - Get Account (XML)
* [GetAccountJSON](docs/sdks/users/README.md#getaccountjson) - Get Account (JSON)
* [AcceptInvite](docs/sdks/users/README.md#acceptinvite) - Accept an Invite
* [DeleteHomeUser](docs/sdks/users/README.md#deletehomeuser) - Delete Home User
* [UpdateHomeUser](docs/sdks/users/README.md#updatehomeuser) - Update Home User
* [UpdateRestrictedUser](docs/sdks/users/README.md#updaterestricteduser) - Update Restricted User
* [GetServerDetails](docs/sdks/users/README.md#getserverdetails) - Get Server Details
* [ShareServerLegacy](docs/sdks/users/README.md#shareserverlegacy) - Share Server (Legacy v1)
* [RemoveShare](docs/sdks/users/README.md#removeshare) - Remove Share
* [UpdateShare](docs/sdks/users/README.md#updateshare) - Update Share
* [GetUserOptOuts](docs/sdks/users/README.md#getuseroptouts) - Get User Opt-Outs

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo/models/operations"
	"github.com/LukeHagar/plexgo/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithAccepts(components.AcceptsApplicationXML),
		plexgo.WithClientIdentifier("abc123"),
		plexgo.WithProduct("Plex for Roku"),
		plexgo.WithVersion("2.4.1"),
		plexgo.WithPlatform("Roku"),
		plexgo.WithPlatformVersion("4.3 build 1057"),
		plexgo.WithDevice("Roku 3"),
		plexgo.WithModel("4200X"),
		plexgo.WithDeviceVendor("Roku"),
		plexgo.WithDeviceName("Living Room TV"),
		plexgo.WithMarketplace("googlePlay"),
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.General.GetServerInfo(ctx, operations.GetServerInfoRequest{}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo/models/operations"
	"github.com/LukeHagar/plexgo/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		plexgo.WithAccepts(components.AcceptsApplicationXML),
		plexgo.WithClientIdentifier("abc123"),
		plexgo.WithProduct("Plex for Roku"),
		plexgo.WithVersion("2.4.1"),
		plexgo.WithPlatform("Roku"),
		plexgo.WithPlatformVersion("4.3 build 1057"),
		plexgo.WithDevice("Roku 3"),
		plexgo.WithModel("4200X"),
		plexgo.WithDeviceVendor("Roku"),
		plexgo.WithDeviceName("Living Room TV"),
		plexgo.WithMarketplace("googlePlay"),
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.General.GetServerInfo(ctx, operations.GetServerInfoRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetServerInfo` function may return the following errors:

| Error Type         | Status Code | Content Type     |
| ------------------ | ----------- | ---------------- |
| sdkerrors.Error    | 401         | application/json |
| sdkerrors.SDKError | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo/models/operations"
	"github.com/LukeHagar/plexgo/models/sdkerrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithAccepts(components.AcceptsApplicationXML),
		plexgo.WithClientIdentifier("abc123"),
		plexgo.WithProduct("Plex for Roku"),
		plexgo.WithVersion("2.4.1"),
		plexgo.WithPlatform("Roku"),
		plexgo.WithPlatformVersion("4.3 build 1057"),
		plexgo.WithDevice("Roku 3"),
		plexgo.WithModel("4200X"),
		plexgo.WithDeviceVendor("Roku"),
		plexgo.WithDeviceName("Living Room TV"),
		plexgo.WithMarketplace("googlePlay"),
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.General.GetServerInfo(ctx, operations.GetServerInfoRequest{})
	if err != nil {

		var e *sdkerrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                                                     | Variables                                    | Description |
| --- | ---------------------------------------------------------- | -------------------------------------------- | ----------- |
| 0   | `https://{IP-description}.{identifier}.plex.direct:{port}` | `identifier`<br/>`IP-description`<br/>`port` |             |
| 1   | `{protocol}://{host}:{port}`                               | `host`<br/>`port`<br/>`protocol`             |             |
| 2   | `https://{full_server_url}`                                | `full_server_url`                            |             |

If the selected server has variables, you may override its default values using the associated option(s):

| Variable          | Option                                    | Default                              | Description                                                                                                                                                                                                                                                                                                                                                                    |
| ----------------- | ----------------------------------------- | ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `identifier`      | `WithIdentifier(identifier string)`       | `"0123456789abcdef0123456789abcdef"` | The unique identifier of this particular PMS                                                                                                                                                                                                                                                                                                                                   |
| `IP-description`  | `WithIPDescription(ipDescription string)` | `"1-2-3-4"`                          | A `-` separated string of the IPv4 or IPv6 address components                                                                                                                                                                                                                                                                                                                  |
| `port`            | `WithPort(port string)`                   | `"32400"`                            | The Port number configured on the PMS. Typically (`32400`).<br/>If using a reverse proxy, this would be the port number configured on the proxy.                                                                                                                                                                                                                               |
| `host`            | `WithHost(host string)`                   | `"localhost"`                        | The Host of the PMS.<br/>If using on a local network, this is the internal IP address of the server hosting the PMS.<br/>If using on an external network, this is the external IP address for your network, and requires port forwarding.<br/>If using a reverse proxy, this would be the external DNS domain for your network, and requires the proxy handle port forwarding. |
| `protocol`        | `WithProtocol(protocol string)`           | `"http"`                             | The network protocol to use. Typically (`http` or `https`)                                                                                                                                                                                                                                                                                                                     |
| `full_server_url` | `WithFullServerURL(fullServerURL string)` | `"http://localhost:32400"`           | The full manual URL to access the PMS                                                                                                                                                                                                                                                                                                                                          |

#### Example

```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithServerIndex(0),
		plexgo.WithIdentifier("0123456789abcdef0123456789abcdef"),
		plexgo.WithIPDescription("1-2-3-4"),
		plexgo.WithPort("32400"),
		plexgo.WithAccepts(components.AcceptsApplicationXML),
		plexgo.WithClientIdentifier("abc123"),
		plexgo.WithProduct("Plex for Roku"),
		plexgo.WithVersion("2.4.1"),
		plexgo.WithPlatform("Roku"),
		plexgo.WithPlatformVersion("4.3 build 1057"),
		plexgo.WithDevice("Roku 3"),
		plexgo.WithModel("4200X"),
		plexgo.WithDeviceVendor("Roku"),
		plexgo.WithDeviceName("Living Room TV"),
		plexgo.WithMarketplace("googlePlay"),
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.General.GetServerInfo(ctx, operations.GetServerInfoRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithServerURL("https://http://localhost:32400"),
		plexgo.WithAccepts(components.AcceptsApplicationXML),
		plexgo.WithClientIdentifier("abc123"),
		plexgo.WithProduct("Plex for Roku"),
		plexgo.WithVersion("2.4.1"),
		plexgo.WithPlatform("Roku"),
		plexgo.WithPlatformVersion("4.3 build 1057"),
		plexgo.WithDevice("Roku 3"),
		plexgo.WithModel("4200X"),
		plexgo.WithDeviceVendor("Roku"),
		plexgo.WithDeviceName("Living Room TV"),
		plexgo.WithMarketplace("googlePlay"),
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.General.GetServerInfo(ctx, operations.GetServerInfoRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.General.GetUserWebhooks(ctx, operations.WithServerURL("https://plex.tv/api/v2"))
	if err != nil {
		log.Fatal(err)
	}
	if res.WebhookPayload != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/LukeHagar/plexgo"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = plexgo.New(plexgo.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name    | Type   | Scheme  |
| ------- | ------ | ------- |
| `Token` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithSecurity("<YOUR_API_KEY_HERE>"),
		plexgo.WithAccepts(components.AcceptsApplicationXML),
		plexgo.WithClientIdentifier("abc123"),
		plexgo.WithProduct("Plex for Roku"),
		plexgo.WithVersion("2.4.1"),
		plexgo.WithPlatform("Roku"),
		plexgo.WithPlatformVersion("4.3 build 1057"),
		plexgo.WithDevice("Roku 3"),
		plexgo.WithModel("4200X"),
		plexgo.WithDeviceVendor("Roku"),
		plexgo.WithDeviceName("Living Room TV"),
		plexgo.WithMarketplace("googlePlay"),
	)

	res, err := s.General.GetServerInfo(ctx, operations.GetServerInfoRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/components"
	"github.com/LukeHagar/plexgo/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := plexgo.New(
		plexgo.WithAccepts(components.AcceptsApplicationXML),
		plexgo.WithClientIdentifier("abc123"),
		plexgo.WithProduct("Plex for Roku"),
		plexgo.WithVersion("2.4.1"),
		plexgo.WithPlatform("Roku"),
		plexgo.WithPlatformVersion("4.3 build 1057"),
		plexgo.WithDevice("Roku 3"),
		plexgo.WithModel("4200X"),
		plexgo.WithDeviceVendor("Roku"),
		plexgo.WithDeviceName("Living Room TV"),
		plexgo.WithMarketplace("googlePlay"),
	)

	res, err := s.Authentication.CreateOAuthPin(ctx, operations.CreateOAuthPinRequest{}, operations.CreateOAuthPinSecurity{
		ClientIdentifier: "<YOUR_API_KEY_HERE>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
