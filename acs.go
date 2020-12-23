package census

import (
	"fmt"
	"time"
)
// https://api.census.gov/data/2019/acs/acs1/subject?get=NAME,S0101_C01_001E&for=new%20england%20city%20and%20town%20area:71650
func ACS(year int, interval int8, f *Form) ([]byte, error) {
	if year < 2005 || year > time.Now().Year() {
		return nil, fmt.Errorf("year must be greater than 2005 and not in the future, got %v", year)
	}

	if interval != 1 && interval != 3 && interval != 5 {
		return nil, fmt.Errorf("acs surveys are 1, 3, and 5 year intervals, but got %v for interval argument", interval)
	}

	return Query(fmt.Sprintf("/%v/acs/acs%v", year, interval), f.toValues())
}
