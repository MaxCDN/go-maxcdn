package maxcdn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Here lies known endpoints.
const (
	PopularFilesEndpoint = "/reports/popularfiles.json"
	StatsEndpoint        = "/reports/stats.json"
)

// GenericResponse is the generic data type for JSON responses from API calls.
type GenericResponse struct {
	Code  int                    `json:"code"`
	Data  map[string]interface{} `json:"data"`
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
	Raw      []byte         // include raw json in GenericResponse
	Response *http.Response // include response in GenericResponse
}

// Parse turns an http response in to a GenericResponse
func (mapper *GenericResponse) Parse(raw []byte) (err error) {
	err = json.Unmarshal(raw, &mapper)
	if err != nil {
		return err
	}

	if mapper.Error.Message != "" || mapper.Error.Type != "" {
		err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

	return err
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
// Generated using http://mervine.net/json2struct
// - changed float64 values to int
type PopularFiles struct {
	Code int `json:"code"`
	Data struct {
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
	} `json:"data"`

	// Added for extra support, see maxcdn.GenericResponse
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
	Raw      []byte
	Response *http.Response
}

// Parse turns an http response in to a PopularFiles
func (mapper *PopularFiles) Parse(raw []byte) (err error) {
	mapper.Raw = raw

	err = json.Unmarshal(raw, &mapper)
	if err != nil {
		return err
	}

	if mapper.Error.Message != "" || mapper.Error.Type != "" {
		err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

	return err
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
	Code int `json:"code"`
	Data struct {
		Stats Stats `json:"stats"`
		//Stats struct {
		//CacheHit    string `json:"cache_hit"`
		//Hit         string `json:"hit"`
		//NoncacheHit string `json:"noncache_hit"`
		//Size        string `json:"size"`
		//Timestamp   string `json:"timestamp"`
		//} `json:"stats"`
		Total string `json:"total"`
	} `json:"data"`
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
	Raw      []byte
	Response *http.Response
}

// Parse turns an http response in to a StatsSummary
func (mapper *SummaryStats) Parse(raw []byte) (err error) {
	mapper.Raw = raw

	err = json.Unmarshal(raw, &mapper)
	if err != nil {
		return err
	}

	if mapper.Error.Message != "" || mapper.Error.Type != "" {
		err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

	return err
}

// StatsSummary is the mapper for /reports/stats.json/{report_type}
// Generated using http://mervine.net/json2struct
// - changed float64 values to int
// - modified Stats and Summary to use Stats definition above
type MultiStats struct {
	Code int `json:"code"`
	Data struct {
		CurrentPageSize int     `json:"current_page_size"`
		Page            int     `json:"page"`
		PageSize        string  `json:"page_size"`
		Pages           int     `json:"pages"`
		Stats           []Stats `json:"stats"`
		Summary         Stats   `json:"summary"`
		Total           string  `json:"total"`
	} `json:"data"`
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
	Raw      []byte
	Response *http.Response
}

// Parse turns an http response in to a Stats
func (mapper *MultiStats) Parse(raw []byte) (err error) {
	mapper.Raw = raw

	err = json.Unmarshal(raw, &mapper)
	if err != nil {
		return err
	}

	if mapper.Error.Message != "" || mapper.Error.Type != "" {
		err = fmt.Errorf("%s: %s", mapper.Error.Type, mapper.Error.Message)
	}

	return err
}
