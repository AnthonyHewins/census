package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AnthonyHewins/census"
)

func acsQuery(yearStr, intervalStr, jsonFile string) {
	year := strToInt(yearStr)
	interval := parseInterval(intervalStr )
	buf := readFile(jsonFile)

	var f census.Form
	if err := json.Unmarshal(buf, &f); err != nil {
		fmt.Println("You passed invalid JSON:")
		fmt.Println(string(buf))
		os.Exit(1)
	}

	resp, err := census.ACS(year, interval, &f)
	if err != nil {
		fmt.Println("Error querying the census API:")
		fmt.Println(err)
		os.Exit(1)
	}

	printAsJson(resp)
}

func parseInterval(interval string) census.ACSInterval {
	switch interval {
	case "se", "supplemental":
		return census.ACSOneYearSupp
	default:
		return census.ACSInterval(strToInt(interval))
	}
}
