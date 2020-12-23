package census

import (
	"fmt"
	"time"
)

func ACS(year int, interval int8, f *Form) ([][]interface{}, error) {
	if year < 2005 || year > time.Now().Year() {
		return nil, fmt.Errorf("year must be greater than 2005 and not in the future, got %v", year)
	}

	if interval != 1 && interval != 3 && interval != 5 {
		return nil, fmt.Errorf("acs surveys are 1, 3, and 5 year intervals, but got %v for interval argument", interval)
	}

	return Query(fmt.Sprintf("/%v/acs/acs%v", year, interval), f.toValues())
}
