package main

import (
	"fmt"
	"os"

)

func main() {
	n := len(os.Args)

	switch n {
	case 2:
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help" {
			help(0)
		}

		genericQuery(os.Args[1], "")
	case 3:
		genericQuery(os.Args[1], os.Args[2])
	case 5:
		if os.Args[1] != "acs" {
			help(1, "dont understand command:", os.Args[1])
		}

		acsQuery(os.Args[2], os.Args[3], os.Args[4])
	default:
		help(1, "wrong number of args:", n)
	}
}

const helpText = `usage: census COMMAND
where COMMAND expands to one of the following, matching best as possible first by how many args, then by command matching, in decreasing priority:

  (-h|--help|help)              Print this help message

  (acs|a) YEAR INTERVAL FILE    Get ACS data for YEAR for INTERVAL (e.g. 2015 for year conducted, 1 to denote
                                ACS 1-year study) and use FILE (a JSON file) to use as a payload. JSON should
                                be structured like below. Different fields are required depending on the endpoint.

                                {
                                    // The key to use. Alternatively, use $CENSUS_API_KEY
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
                                    // Since the census API is complex, if this input is a string and cannot be parsed
                                    // it will just be passed off to the census API in case you have a complex query you
                                    // want to use.
                                    "for": 0,
                                    "in": "", // restrict at areas smaller than state level

                                    // Require data is exactly on this time (follow this formatting).
                                    // "year" is required, but "month" is not
                                    // Mutually exclusive from startTime/endTime
                                    "onTime": {"year": 2015, "month": 2},

                                    // Require data is on or after this time
                                    "startTime": {"year": 2014},
                                    // Require data is on or before this time
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
