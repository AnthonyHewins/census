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
	Year int
	Month time.Month
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
	Get []string

	// Predicates for the query. Slices are valid
	Predicate map[string]interface{}

	// Geography fields
	For string
	In string

	// Time Fields
	OnTime *Time
	StartTime *Time
	EndTime *Time
}

func (f Form) toValues() url.Values {
	values := url.Values{
		"get": {strings.Join(f.Get, ",")},
		"in": {f.In},
		"for": {f.For},
	}

	if f.OnTime != nil {
		values.Add("time", f.OnTime.String())
	} else {
		timeQuery := []string{}

		if f.StartTime != nil {
			timeQuery = append(timeQuery, fmt.Sprintf("from+%v", f.StartTime))
		}

		if f.EndTime != nil {
			timeQuery = append(timeQuery, fmt.Sprintf("to+%v", f.EndTime))
		}

		values.Add("time", strings.Join(timeQuery, "+"))
	}

	for field, predicate := range f.Predicate {
		switch p := predicate.(type) {
		case []interface{}:
			for _, condition := p {
				values.Add(field, fmt.Sprint(condition))
			}
		case interface{}:
			values.Add(field, fmt.Sprint(p))
		}
	}
}

// Query is a generic function for handling generic queries that may not be handled by other
// endpoint functions. `path` will be appended to the base URL (`api.census.gov/data`) and
// `params` will be encoded into GET params
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
