package maxcdn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GenericResponse is the generic data type for JSON responses from API calls.
type GenericResponse struct {
	Code  float64                `json:"code"`
	Data  map[string]interface{} `json:"data"`
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
	Raw string // include raw json in GenericResponse
}

// Parse turns an http response in to a GenericResponse
func (mapper *GenericResponse) Parse(r *http.Response) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &mapper)

	mapper.Raw = string(data) // include raw json in GenericResponse
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
// - /reports/stats.json/{report_type}
// - /reports/popularfiles.json
