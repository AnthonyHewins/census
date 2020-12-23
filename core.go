package census

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"net/http"
)

const endpoint = "https://api.census.gov/data"

func Query(path string, params url.Values) ([]byte, error) {
	resp, err := http.Get(
		fmt.Sprintf("%v%v?%v", endpoint, path, params.Encode()),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
