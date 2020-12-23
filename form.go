package census

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

var censusApiKey = os.Getenv("CENSUS_API_KEY")

// Time is a lightweight struct for specifying
// a month and year
type Time struct {
	Year int `json:"year"`
	Month time.Month `json:"month"`
}

func (t Time) String() string {
	if t.Month != 0 {
		return fmt.Sprintf("%v-%v", t.Year, t.Month)
	}

	return fmt.Sprint(t.Year)
}

type Form struct {
	Key string `json:"key"`

	// What fields to get
	Get []string `json:"get"`

	// Predicates for the query. Slices are valid
	Predicate map[string]interface{} `json:"predicates"`

	// Geography fields
	For int `json:"for"` // restrict at various levels
	In int `json:"in"` // restrict at areas smaller than state level, after already using "for"

	// Time Fields
	OnTime *Time `json:"onTime"`
	StartTime *Time `json:"startTime"`
	EndTime *Time `json:"endTime"`
}

func (f Form) toValues() url.Values {
	values := url.Values{
		"get": {strings.Join(f.Get, ",")},
	}

	if f.Key != "" {
		values.Add("key", f.Key)
	} else if censusApiKey != "" {
		values.Add("key", censusApiKey)
	}

	if f.In != 0 {
		values.Add("in", fmt.Sprint(f.In))
	}

	if f.For != 0 {
		values.Add("for", fmt.Sprint(f.For))
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
