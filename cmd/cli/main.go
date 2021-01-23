package main

import (
	"fmt"
	"os"
)

func main() {
	n := len(os.Args)

	switch n {
	case 2:
		switch os.Args[1] {
		case "-h", "--help", "help", "h":
			help(0)
		default:
			help(1, "dont understand command:", os.Args[1])
		}
	case 5:
		switch os.Args[1] {
		case "acs", "a":
			acsQuery(os.Args[2], os.Args[3], os.Args[4])
		default:
			help(1, "dont understand command:", os.Args[1])
		}
	default:
		help(1, "wrong number of args:", n)
	}
}

const helpText = `usage: census COMMAND
where COMMAND expands to one of the following, matching best as possible first by how many args, then by command matching, in decreasing priority:

  (-h| --help | help | h)       Print this help message

  (acs|a) YEAR INTERVAL FILE    Get ACS data collected on YEAR for INTERVAL (e.g. 2015 for year conducted, 1 to denote
                                ACS 1-year study) and use FILE (a JSON file) to use as a payload. JSON should
                                be structured like below. Different fields are required depending on the endpoint.

                                INTERVAL is one of 1, 3, 5, or "se"

                                {
                                	// The key to use. Alternatively, $CENSUS_API_KEY will be checked for a fallback value
                                	"key": "your-api-key",

                                	// What fields to get, a string array
                                	"get": ["NAME"],

                                	// predicates to filter on
                                	"predicate": {
                                		"NAME": [1, 2] // NAME must equal 1 or 2
                                		"FIELD": 12    // FIELD must equal 12
                                	},

                                	// Geography fields. You'll want to use FIPS codes:
                                	// https://en.wikipedia.org/wiki/Federal_Information_Processing_Standard_state_code
                                	"for": 0,
                                	"in": "", // restrict at areas smaller than state level

                                	// Filter results that occur on this time
                                	// "year" is required, but "month" is not
                                	// Mutually exclusive from startTime/endTime
                                	"onTime": {"year": 2015, "month": 2},

                                	// Filter results on or after this time
                                	"startTime": {"year": 2014},
                                	// Filter results on or before this time
                                	"endTime": {"year": 2016},
                                }

  PATH [JSON]                   The most generalized query: Make a request to PATH and use the string JSON
                                as the URL-encoded form for the request. Leaving JSON blank is equivalent to having no form
`

func help(exitCode int, extraMessages ...interface{}) {
	fmt.Println(extraMessages...)
	fmt.Println(helpText)
	os.Exit(exitCode)
}
