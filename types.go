package maxcdn

import (
	"encoding/json"
	"net/http"
)

// Response object for all json requests.
type Response struct {
	Code  int             `json:"code,omitempty"`
	Data  json.RawMessage `json:"data,omitempty"`
	Error struct {
		Message string `json:"message,omitempty"`
		Type    string `json:"type,omitempty"`
	} `json:"error,omitempty"`

	// Non-JSON data.
	Headers *http.Header
}

// Generic is the generic data type for JSON responses from API calls.
type Generic struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

// Account is for /account.json
type Account struct {
	Account struct {
		Alias                  string `json:"alias,omitempty"`
		DateCreated            string `json:"date_created,omitempty"`
		DateUpdated            string `json:"date_updated,omitempty"`
		DefaultPullZoneIpID    string `json:"default_pull_zone_ip_id,omitempty"`
		DefaultPushZoneIpID    string `json:"default_push_zone_ip_id,omitempty"`
		DefaultStorageIpID     string `json:"default_storage_ip_id,omitempty"`
		DefaultVodDirectIpID   string `json:"default_vod_direct_ip_id,omitempty"`
		DefaultVodPseudoIpID   string `json:"default_vod_pseudo_ip_id,omitempty"`
		DefaultVodRtmpIpID     string `json:"default_vod_rtmp_ip_id,omitempty"`
		DefaultVodStorageIpID  string `json:"default_vod_storage_ip_id,omitempty"`
		EdgerulesCredits       string `json:"edgerules_credits,omitempty"`
		FlexCredits            string `json:"flex_credits,omitempty"`
		ID                     string `json:"id,omitempty"`
		Name                   string `json:"name,omitempty"`
		SecureTokenPullCredits string `json:"secure_token_pull_credits,omitempty"`
		SslCredits             string `json:"ssl_credits,omitempty"`
		Status                 string `json:"status,omitempty"`
		StorageQuota           string `json:"storage_quota,omitempty"`
		ZoneCredits            string `json:"zone_credits,omitempty"`
	} `json:"account,omitempty"`
}

type AccountAddress struct {
	Address struct {
		City        string `json:"city,omitempty"`
		Country     string `json:"country,omitempty"`
		DateCreated string `json:"date_created,omitempty"`
		DateUpdated string `json:"date_updated,omitempty"`
		ID          string `json:"id,omitempty"`
		State       string `json:"state,omitempty"`
		Street1     string `json:"street1,omitempty"`
		Street2     string `json:"street2,omitempty"`
		Zip         string `json:"zip,omitempty"`
	} `json:"address,omitempty"`
}

// PopularFiles is the mapper for /reports/popularfiles.json
//
// Maps to 'Endpoint.Reports.PopularFiles'
type PopularFiles struct {
	CurrentPageSize int    `json:"current_page_size,omitempty"`
	Page            int    `json:"page,omitempty"`
	PageSize        string `json:"page_size,omitempty"`
	Pages           int    `json:"pages,omitempty"`
	PopularFiles    []struct {
		BucketID  string `json:"bucket_id,omitempty"`
		Hit       string `json:"hit,omitempty"`
		Size      string `json:"size,omitempty"`
		Timestamp string `json:"timestamp,omitempty"`
		Uri       string `json:"uri,omitempty"`
		Vhost     string `json:"vhost,omitempty"`
	} `json:"popularfiles,omitempty"`
	Summary struct {
		Hit  string `json:"hit,omitempty"`
		Size string `json:"size,omitempty"`
	} `json:"summary,omitempty"`
	Total string `json:"total,omitempty"`
}

// StatsSummary is the mapper for /reports/stats.json
//
// Maps to 'Endpoint.Reports.Stats'
type StatsSummary struct {
	Stats struct {
		CacheHit    string `json:"cache_hit,omitempty"`
		Hit         string `json:"hit,omitempty"`
		NoncacheHit string `json:"noncache_hit,omitempty"`
		Size        string `json:"size,omitempty"`
	} `json:"stats,omitempty"`
	Total string `json:"total,omitempty"`
}

// StatsSummary is the mapper for /reports/stats.json/{report_type}
//
// Maps to 'Endpoint.Reports.StatsBy("{report_type}")'
type Stats struct {
	CurrentPageSize int    `json:"current_page_size,omitempty"`
	Page            int    `json:"page,omitempty"`
	PageSize        string `json:"page_size,omitempty"`
	Pages           int    `json:"pages,omitempty"`
	Stats           []struct {
		CacheHit    string `json:"cache_hit,omitempty"`
		Hit         string `json:"hit,omitempty"`
		NoncacheHit string `json:"noncache_hit,omitempty"`
		Size        string `json:"size,omitempty"`
		Timestamp   string `json:"timestamp,omitempty"`
	} `json:"stats,omitempty"`
	Summary struct {
		Stats struct {
			CacheHit    string `json:"cache_hit,omitempty"`
			Hit         string `json:"hit,omitempty"`
			NoncacheHit string `json:"noncache_hit,omitempty"`
			Size        string `json:"size,omitempty"`
			Timestamp   string `json:"timestamp,omitempty"`
		} `json:"stats,omitempty"`
		Total string `json:"total,omitempty"`
	} `json:"summary,omitempty"`
	Total string `json:"total,omitempty"`
}

// Pullzone is for POST|PUT /zones/pull.json and GET /zones/pull.json/{zone_id}
//
// Maps to 'Endpoint.Zones.Pull' and 'Endpoint.Zones.PullBy("{zone_id}")
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
	} `json:"pullzone,omitempty"`
}

// Pullzones is for GET /zones/pull.json
//
// Maps to 'Endpoint.Zones.Pull'
type Pullzones struct {
	CurrentPageSize int    `json:"current_page_size,omitempty"`
	Page            int    `json:"page,omitempty"`
	PageSize        string `json:"page_size,omitempty"`
	Pages           int    `json:"pages,omitempty"`
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
	} `json:"pullzones,omitempty"`
	Total int `json:"total,omitempty"`
}

// Users is for GET /users.json
type Users struct {
	CurrentPageSize int    `json:"current_page_size,omitempty"`
	Page            int    `json:"page,omitempty"`
	PageSize        string `json:"page_size,omitempty"`
	Pages           int    `json:"pages,omitempty"`
	Total           int    `json:"total,omitempty"`
	Users           []struct {
		BrandID          string   `json:"brand_id,omitempty"`
		DateCreated      string   `json:"date_created,omitempty"`
		DateLastLogin    string   `json:"date_last_login,omitempty"`
		DateUpdated      string   `json:"date_updated,omitempty"`
		DefaultCompanyID string   `json:"default_company_id,omitempty"`
		Email            string   `json:"email,omitempty"`
		Firstname        string   `json:"firstname,omitempty"`
		ID               string   `json:"id,omitempty"`
		IpLastLogin      string   `json:"ip_last_login,omitempty"`
		IsAdmin          string   `json:"isadmin,omitempty"`
		Isdisabled       string   `json:"isdisabled,omitempty"`
		Lastname         string   `json:"lastname,omitempty"`
		LoginWhitelist   string   `json:"login_whitelist,omitempty"`
		Phone            string   `json:"phone,omitempty"`
		Roles            []string `json:"roles,omitempty"`
		Timezone         string   `json:"timezone,omitempty"`
	} `json:"users,omitempty"`
}

// User is for GET /users.json/{user_id}
type User struct {
	User struct {
		BrandID          string   `json:"brand_id,omitempty"`
		DateCreated      string   `json:"date_created,omitempty"`
		DateLastLogin    string   `json:"date_last_login,omitempty"`
		DateUpdated      string   `json:"date_updated,omitempty"`
		DefaultCompanyID string   `json:"default_company_id,omitempty"`
		Email            string   `json:"email,omitempty"`
		Firstname        string   `json:"firstname,omitempty"`
		ID               string   `json:"id,omitempty"`
		IpLastLogin      string   `json:"ip_last_login,omitempty"`
		IsAdmin          string   `json:"isadmin,omitempty"`
		Isdisabled       string   `json:"isdisabled,omitempty"`
		Lastname         string   `json:"lastname,omitempty"`
		LoginWhitelist   string   `json:"login_whitelist,omitempty"`
		Phone            string   `json:"phone,omitempty"`
		Roles            []string `json:"roles,omitempty"`
		Timezone         string   `json:"timezone,omitempty"`
	} `json:"user,omitempty"`
}
