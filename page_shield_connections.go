package cloudflare

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
)

// ListPageShieldConnectionsParams represents parameters for a page shield connection request
type ListPageShieldConnectionsParams struct {
	Direction           string `json:"direction"`
	ExcludeCdnCgi       bool   `json:"exclude_cdn_cgi"`
	ExcludeUrls         string `json:"exclude_urls"`
	Export              string `json:"export"`
	Hosts               string `json:"hosts"`
	OrderBy             string `json:"order_by"`
	Page                string `json:"page"`
	PageURL             string `json:"page_url"`
	PerPage             int    `json:"per_page"`
	PrioritizeMalicious bool   `json:"prioritize_malicious"`
	Status              string `json:"status"`
	URLs                string `json:"urls"`
}

// PageShieldConnection represents a page shield connection
type PageShieldConnection struct {
	AddedAt                 string   `json:"added_at"`
	DomainReportedMalicious bool     `json:"domain_reported_malicious"`
	FirstPageURL            string   `json:"first_page_url"`
	FirstSeenAt             string   `json:"first_seen_at"`
	Host                    string   `json:"host"`
	ID                      string   `json:"id"`
	LastSeenAt              string   `json:"last_seen_at"`
	PageURLs                []string `json:"page_urls"`
	URL                     string   `json:"url"`
	URLContainsCdnCgiPath   bool     `json:"url_contains_cdn_cgi_path"`
}

// ListPageShieldConnectionsResponse represents the response from the list page shield connections endpoint
type ListPageShieldConnectionsResponse struct {
	Result []PageShieldConnection `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// ListPageShieldConnections lists all page shield connections for a zone
//
// API documentation: https://developers.cloudflare.com/api/operations/page-shield-list-page-shield-connections
func (api *API) ListPageShieldConnections(ctx context.Context, rc *ResourceContainer, params ListPageShieldConnectionsParams) ([]PageShieldConnection, ResultInfo, error) {
	path := fmt.Sprintf("/zones/%s/page_shield/connections", rc.Identifier)

	uri := buildURI(path, params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, ResultInfo{}, err
	}

	var psResponse ListPageShieldConnectionsResponse
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return psResponse.Result, psResponse.ResultInfo, nil
}

// GetPageShieldConnection gets a page shield connection for a zone
//
// API documentation: https://developers.cloudflare.com/api/operations/page-shield-get-a-page-shield-connection
func (api *API) GetPageShieldConnection(ctx context.Context, rc *ResourceContainer, connectionID string) (*PageShieldConnection, error) {
	path := fmt.Sprintf("/zones/%s/page_shield/connections/%s", rc.Identifier, connectionID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var psResponse PageShieldConnection
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &psResponse, nil
}
