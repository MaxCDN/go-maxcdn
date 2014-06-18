package maxcdn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type reports struct {
	PopularFiles string
	Stats        string
}

func (r *reports) StatsBy(t string) string {
	return fmt.Sprintf("%s/%s", r.Stats, t)
}

type zones struct {
	Pull string
	Push string
}

func (z *zones) PullBy(t int) string {
	return fmt.Sprintf("%s/%d", z.Pull, t)
}

func (z *zones) PullByString(t string) string {
	return fmt.Sprintf("%s/%s", z.Pull, t)
}

func (z *zones) PullCacheBy(t int) string {
	return fmt.Sprintf("%s/%d/cache", z.Pull, t)
}

func (z *zones) PullCacheByString(t string) string {
	return fmt.Sprintf("%s/%s/cache", z.Pull, t)
}

func (z *zones) PushBy(t int) string {
	return fmt.Sprintf("%s/%d", z.Push, t)
}

func (z *zones) PushByString(t string) string {
	return fmt.Sprintf("%s/%s", z.Push, t)
}

type endpoints struct {
	Account        string
	AccountAddress string
	Reports        *reports
	Zones          *zones
}

/* Endpoint usage examples:
 *
 *  // for pull zone with id of '123456'
 *  e := Endpoint.Zones.PullBy('123456')
 *  => /zones/pull.json/123456
 *
 *  // for popular files report
 *  e := Endpoint.Reports.PopularFiles
 *  => /reports/popularfiles.json
 *
 *  // for hourly stats report
 *  e := Endpoint.Reports.StatsBy('hourly')
 *  => /reports/stats.json/hourly
 */
var Endpoint = endpoints{
	Account:        "/account.json",
	AccountAddress: "/account.json/address",
	Reports: &reports{
		PopularFiles: "/reports/popularfiles.json",
		Stats:        "/reports/stats.json",
	},
	Zones: &zones{
		Pull: "/zones/pull.json",
		Push: "/zones/push.json",
	},
}

// Response object for all json requests.
type Response struct {
	Code  int             `json:"code"`
	Data  json.RawMessage `json:"data"`
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`

	// Non-JSON data.
	Headers *http.Header
}

// Generic is the generic data type for JSON responses from API calls.
type Generic struct {
	Data map[string]interface{} `json:"data"`
}

// Account is for /account.json
type Account struct {
	Account struct {
		Alias                  string `json:"alias"`
		DateCreated            string `json:"date_created"`
		DateUpdated            string `json:"date_updated"`
		DefaultPullZoneIpID    string `json:"default_pull_zone_ip_id"`
		DefaultPushZoneIpID    string `json:"default_push_zone_ip_id"`
		DefaultStorageIpID     string `json:"default_storage_ip_id"`
		DefaultVodDirectIpID   string `json:"default_vod_direct_ip_id"`
		DefaultVodPseudoIpID   string `json:"default_vod_pseudo_ip_id"`
		DefaultVodRtmpIpID     string `json:"default_vod_rtmp_ip_id"`
		DefaultVodStorageIpID  string `json:"default_vod_storage_ip_id"`
		EdgerulesCredits       string `json:"edgerules_credits"`
		FlexCredits            string `json:"flex_credits"`
		ID                     string `json:"id"`
		Name                   string `json:"name"`
		SecureTokenPullCredits string `json:"secure_token_pull_credits"`
		SslCredits             string `json:"ssl_credits"`
		Status                 string `json:"status"`
		StorageQuota           string `json:"storage_quota"`
		ZoneCredits            string `json:"zone_credits"`
	} `json:"account"`
}

type AccountAddress struct {
	Address struct {
		City        string `json:"city"`
		Country     string `json:"country"`
		DateCreated string `json:"date_created"`
		DateUpdated string `json:"date_updated"`
		ID          string `json:"id"`
		State       string `json:"state"`
		Street1     string `json:"street1"`
		Street2     string `json:"street2"`
		Zip         string `json:"zip"`
	} `json:"address"`
}

// Specific types
//
// TODO:
// Add specific types for more commonly called endpoints like:
//
// - /account.json
// - /account.json/address
// - /users.json
// - /users.json/{user_id}
// - /zones.json
// - /zones/pull.json
// - /zones/push.json

/*
 * Adding custom mapping below as needed.
 */

// PopularFiles is the mapper for /reports/popularfiles.json
type PopularFiles struct {
	CurrentPageSize int    `json:"current_page_size"`
	Page            int    `json:"page"`
	PageSize        string `json:"page_size"`
	Pages           int    `json:"pages"`
	Popularfiles    []struct {
		BucketID  string `json:"bucket_id"`
		Hit       string `json:"hit"`
		Size      string `json:"size"`
		Timestamp string `json:"timestamp"`
		Uri       string `json:"uri"`
		Vhost     string `json:"vhost"`
	} `json:"popularfiles"`
	Summary struct {
		Hit  string `json:"hit"`
		Size string `json:"size"`
	} `json:"summary"`
	Total string `json:"total"`
}

// Stats is to be used within MultiStats and SummaryStats to hold
// the core stats data.
type Stats struct {
	CacheHit    string `json:"cache_hit"`
	Hit         string `json:"hit"`
	NoncacheHit string `json:"noncache_hit"`
	Size        string `json:"size"`
	Timestamp   string `json:"timestamp"`
}

type SummaryStats struct {
	Stats Stats  `json:"stats"`
	Total string `json:"total"`
}

// StatsSummary is the mapper for /reports/stats.json/{report_type}
type MultiStats struct {
	CurrentPageSize int     `json:"current_page_size"`
	Page            int     `json:"page"`
	PageSize        string  `json:"page_size"`
	Pages           int     `json:"pages"`
	Stats           []Stats `json:"stats"`
	Summary         Stats   `json:"summary"`
	Total           string  `json:"total"`
}

// Pullzone is for POST|PUT /zones/pull.json and GET /zones/pull.json/{zone_id}
type Pullzone struct {
	Pullzone struct {
		BackendCompress       string `json:"backend_compress,omitempty"`
		CacheValid            string `json:"cache_valid,omitempty"`
		CanonicalLinkHeaders  string `json:"canonical_link_headers,omitempty"`
		CdnURL                string `json:"cdn_url,omitempty"`
		Compress              string `json:"compress,omitempty"`
		ContentDisposition    string `json:"content_disposition,omitempty"`
		CreationDate          string `json:"creation_date,omitempty"`
		DisallowRobots        string `json:"disallow_robots,omitempty"`
		DisallowRobotsTxt     string `json:"disallow_robots_txt,omitempty"`
		DnsCheck              string `json:"dns_check,omitempty"`
		Expires               string `json:"expires,omitempty"`
		HideSetcookieHeader   string `json:"hide_setcookie_header,omitempty"`
		ID                    string `json:"id,omitempty"`
		IgnoreCacheControl    string `json:"ignore_cache_control,omitempty"`
		IgnoreExpiresHeader   string `json:"ignore_expires_header,omitempty"`
		IgnoreSetcookieHeader string `json:"ignore_setcookie_header,omitempty"`
		Inactive              string `json:"inactive,omitempty"`
		Ip                    string `json:"ip,omitempty"`
		Label                 string `json:"label,omitempty"`
		Locked                string `json:"locked,omitempty"`
		Name                  string `json:"name,omitempty"`
		Port                  string `json:"port,omitempty"`
		ProxyCacheLock        string `json:"proxy_cache_lock,omitempty"`
		ProxyCacheLockTimeout string `json:"proxy_cache_lock_timeout,omitempty"`
		ProxyInactive         string `json:"proxy_inactive,omitempty"`
		PseudoStreaming       string `json:"pseudo_streaming,omitempty"`
		Queries               string `json:"queries,omitempty"`
		SetHostHeader         string `json:"set_host_header,omitempty"`
		Spdy                  string `json:"spdy,omitempty"`
		SpdyHeadersComp       string `json:"spdy_headers_comp,omitempty"`
		Sslshared             string `json:"sslshared,omitempty"`
		Suspend               string `json:"suspend,omitempty"`
		ThrottleFcc           string `json:"throttle_fcc,omitempty"`
		TmpURL                string `json:"tmp_url,omitempty"`
		Type                  string `json:"type,omitempty"`
		UpstreamEnabled       string `json:"upstream_enabled,omitempty"`
		URL                   string `json:"url,omitempty"`
		UseStale              string `json:"use_stale,omitempty"`
		ValidReferers         string `json:"valid_referers,omitempty"`
		WebpEnabled           string `json:"webp_enabled,omitempty"`
		XForwardFor           string `json:"x_forward_for,omitempty"`
	} `json:"pullzone"`
}

// Pullzones is for GET /zones/pull.json
type Pullzones struct {
	CurrentPageSize int    `json:"current_page_size"`
	Page            int    `json:"page"`
	PageSize        string `json:"page_size"`
	Pages           int    `json:"pages"`
	Pullzones       []struct {
		BackendCompress       string `json:"backend_compress,omitempty"`
		CacheValid            string `json:"cache_valid,omitempty"`
		CanonicalLinkHeaders  string `json:"canonical_link_headers,omitempty"`
		CdnURL                string `json:"cdn_url,omitempty"`
		Compress              string `json:"compress,omitempty"`
		ContentDisposition    string `json:"content_disposition,omitempty"`
		CreationDate          string `json:"creation_date,omitempty"`
		DisallowRobots        string `json:"disallow_robots,omitempty"`
		DisallowRobotsTxt     string `json:"disallow_robots_txt,omitempty"`
		DnsCheck              string `json:"dns_check,omitempty"`
		Expires               string `json:"expires,omitempty"`
		HideSetcookieHeader   string `json:"hide_setcookie_header,omitempty"`
		ID                    string `json:"id,omitempty"`
		IgnoreCacheControl    string `json:"ignore_cache_control,omitempty"`
		IgnoreExpiresHeader   string `json:"ignore_expires_header,omitempty"`
		IgnoreSetcookieHeader string `json:"ignore_setcookie_header,omitempty"`
		Inactive              string `json:"inactive,omitempty"`
		Ip                    string `json:"ip,omitempty"`
		Label                 string `json:"label,omitempty"`
		Locked                string `json:"locked,omitempty"`
		Name                  string `json:"name,omitempty"`
		Port                  string `json:"port,omitempty"`
		ProxyCacheLock        string `json:"proxy_cache_lock,omitempty"`
		ProxyCacheLockTimeout string `json:"proxy_cache_lock_timeout,omitempty"`
		ProxyInactive         string `json:"proxy_inactive,omitempty"`
		PseudoStreaming       string `json:"pseudo_streaming,omitempty"`
		Queries               string `json:"queries,omitempty"`
		SetHostHeader         string `json:"set_host_header,omitempty"`
		Spdy                  string `json:"spdy,omitempty"`
		SpdyHeadersComp       string `json:"spdy_headers_comp,omitempty"`
		Sslshared             string `json:"sslshared,omitempty"`
		Suspend               string `json:"suspend,omitempty"`
		ThrottleFcc           string `json:"throttle_fcc,omitempty"`
		TmpURL                string `json:"tmp_url,omitempty"`
		Type                  string `json:"type,omitempty"`
		UpstreamEnabled       string `json:"upstream_enabled,omitempty"`
		URL                   string `json:"url,omitempty"`
		UseStale              string `json:"use_stale,omitempty"`
		ValidReferers         string `json:"valid_referers,omitempty"`
		WebpEnabled           string `json:"webp_enabled,omitempty"`
		XForwardFor           string `json:"x_forward_for,omitempty"`
	} `json:"pullzones"`
	Total int `json:"total"`
}
