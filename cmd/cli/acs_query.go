package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

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
	fmt.Println(string(resp)......)
}

func strToInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		help(1, "argument must be a number, got:", s)
	}

	return int(i)
}

func readFile(name string) []byte {
	buf, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}

	return buf
}
