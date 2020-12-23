package census

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const endpoint = "https://api.census.gov/data"

// Query is a generic function for handling generic queries that may not be handled by other
// endpoint functions. `path` will be appended to the base URL (`api.census.gov/data`) and
// `params` will be encoded into GET params
func Query(path string, params url.Values) ([][]interface{}, error) {
	finalUrl := fmt.Sprintf("%v%v?%v", endpoint, path, params.Encode())

	resp, err := http.Get(finalUrl)

	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var twoDimensional [][]interface{}
	return twoDimensional, json.Unmarshal(buf, &twoDimensional)
}
