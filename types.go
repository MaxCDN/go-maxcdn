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
type Generic map[string]interface{}

// Logs
type Logs struct {
	Limit       int    `json:"limit"`
	NextPageKey string `json:"next_page_key"`
	Page        int    `json:"page"`
	Records     []struct {
		Bytes           int     `json:"bytes"`
		CacheStatus     string  `json:"cache_status"`
		ClientAsn       string  `json:"client_asn"`
		ClientCity      string  `json:"client_city"`
		ClientContinent string  `json:"client_continent"`
		ClientCountry   string  `json:"client_country"`
		ClientDma       string  `json:"client_dma"`
		ClientIp        string  `json:"client_ip"`
		ClientLatitude  float64 `json:"client_latitude"`
		ClientLongitude float64 `json:"client_longitude"`
		ClientState     string  `json:"client_state"`
		CompanyID       int     `json:"company_id"`
		Hostname        string  `json:"hostname"`
		Method          string  `json:"method"`
		OriginTime      float64 `json:"origin_time"`
		Pop             string  `json:"pop"`
		Protocol        string  `json:"protocol"`
		QueryString     string  `json:"query_string"`
		Referer         string  `json:"referer"`
		Scheme          string  `json:"scheme"`
		Status          int     `json:"status"`
		Time            string  `json:"time"`
		Uri             string  `json:"uri"`
		UserAgent       string  `json:"user_agent"`
		ZoneID          int     `json:"zone_id"`
	} `json:"records"`
	RequestTime int `json:"request_time"`
}
