package core

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"net/http"
)

const endpoint = "https://api.census.gov/data"

func Query(path string, params url.Values) (map[string]interface{}, error) {
	resp, err := http.PostForm(fmt.Sprintf("%v%v", endpoint, path), params)

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonObj map[string]interface{}
	return jsonObj, json.Unmarshal(buf, &jsonObj)
}
