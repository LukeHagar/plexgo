/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

package plexgo

import (
	"regexp"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/lukehagar/plexgo/plextv"
	"github.com/lukehagar/plexgo/pms"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIClient manages communication with the Plex API
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	PMS    *pms.APIClient
	PlexTV *plextv.APIClient
	token  string
}

type service struct {
	PMS    *pms.APIClient
	PlexTV *plextv.APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = retryablehttp.NewClient()
	}

	c := &APIClient{}

	ClientPMS := pms.NewConfiguration(cfg.ClientConfiguration.ClientId, cfg.ClientConfiguration.ClientSecret, cfg.ClientConfiguration.BaseURL, cfg.ClientConfiguration.TokenURL, cfg.ClientConfiguration.Token)
	ClientPlexTV := plextv.NewConfiguration(cfg.ClientConfiguration.ClientId, cfg.ClientConfiguration.ClientSecret, "https://plex.tv/api", cfg.ClientConfiguration.TokenURL, cfg.ClientConfiguration.Token)

	ClientPMS.HTTPClient = cfg.HTTPClient
	ClientPlexTV.HTTPClient = cfg.HTTPClient

	c.PMS = pms.NewAPIClient(ClientPMS)
	c.PlexTV = plextv.NewAPIClient(ClientPlexTV)

	// API Services

	return c
}
