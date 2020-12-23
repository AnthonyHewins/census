package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/AnthonyHewins/census"
)

func acsQuery(yearStr, intervalStr, jsonFile string) {
	year := strToInt(yearStr)
	interval := strToInt(intervalStr)
	buf := readFile(jsonFile)

	var f census.Form
	if err := json.Unmarshal(buf, &f); err != nil {
		fmt.Println("You passed invalid JSON:")
		fmt.Println(string(buf))
		os.Exit(1)
	}

	resp, err := census.ACS(year, int8(interval), &f)
	if err != nil {
		log.Fatalln(err)
	}

	// 2D slice interface{}, use ... twice
	printAsJson(resp)
}
