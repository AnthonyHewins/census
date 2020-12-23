package census

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const endpoint = "https://api.census.gov/data"

// Time is a lightweight struct for specifying
// a month and year
type Time struct {
	Year int `json:"year"`
	Month time.Month `json:"month"`
}

func (t Time) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprint(t.Year))

	if t.Month != 0 {
		sb.WriteString("-%v")
		sb.WriteString(t.Month.String())
	}

	return sb.String()
}

type Form struct {
	// What fields to get
	Get []string `json:"get"`

	// Predicates for the query. Slices are valid
	Predicate map[string]interface{} `json:"predicates"`

	// Geography fields
	For string `json:"for"` // restrict at various levels
	In string `json:"in"` // restrict at areas smaller than state level, after already using "for"

	// Time Fields
	OnTime *Time `json:"onTime"`
	StartTime *Time `json:"startTime"`
	EndTime *Time `json:"endTime"`
}

func (f Form) toValues() url.Values {
	values := url.Values{
		"get": {strings.Join(f.Get, ",")},
	}

	if f.In != "" {
		values.Add("in", f.In)
	}

	if f.For != "" {
		values.Add("for", f.For)
	}

	if f.OnTime != nil {
		values.Add("time", f.OnTime.String())
	} else {
		var timeQuery string

		if f.StartTime != nil && f.EndTime != nil {
			timeQuery = fmt.Sprintf("from+%v+to+%v", f.StartTime, f.EndTime)
		} else if f.StartTime != nil {
			timeQuery = fmt.Sprintf("from+%v", f.StartTime)
		} else if f.EndTime != nil {
			timeQuery = fmt.Sprintf("to+%v", f.EndTime)
		}

		if timeQuery != "" {
			values.Add("time", timeQuery)
		}
	}

	for field, predicate := range f.Predicate {
		switch p := predicate.(type) {
		case []interface{}:
			for _, condition := range p {
				values.Add(field, fmt.Sprint(condition))
			}
		default:
			values.Add(field, fmt.Sprint(p))
		}
	}

	return values
}

// Query is a generic function for handling generic queries that may not be handled by other
// endpoint functions. `path` will be appended to the base URL (`api.census.gov/data`) and
// `params` will be encoded into GET params
func Query(path string, params url.Values) ([]byte, error) {
	finalUrl := fmt.Sprintf("%v%v?%v", endpoint, path, params.Encode())
		fmt.Println(finalUrl)
	resp, err := http.Get(
		finalUrl,
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
