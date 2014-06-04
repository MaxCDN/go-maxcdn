package maxcdn

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	//"net/http"
)

// GenericResponse is the generic data type for JSON responses from API calls.
type GenericResponse struct {
	Code  int                    `json:"code"`
	Data  map[string]interface{} `json:"data"`
	Raw   []byte                 // include raw json in GenericResponse
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
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
// - /reports/stats.json/{report_type}
// - /reports/popularfiles.json
