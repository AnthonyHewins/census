package census

import (
	"fmt"
	"time"
)

type ACSInterval int8

const (
	ACSOneYearSupp ACSInterval = -1

	ACSOneYear   ACSInterval = 1
	ACSThreeYear ACSInterval = 3
	ACSFiveYear  ACSInterval = 5
)

// ACS hits one of the American Community Survey endpoints.
// `year` specifies what year, `interval` specifies the interval the survey is taken,
// and `f` is a pointer to the `Form` object.
func ACS(year int, interval ACSInterval, f *Form) ([][]interface{}, error) {
	if year < 2005 || year > time.Now().Year() {
		return nil, fmt.Errorf("year must be greater than 2005 and not in the future, got %v", year)
	}

	var path string
	if interval < 0 {
		path = fmt.Sprintf("/%v/acs/acsse", year)
	} else {
		path = fmt.Sprintf("/%v/acs/acs%v", year, interval)
	}

	return Query(path, f)
}
