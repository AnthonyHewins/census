package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"encoding/json"
)

func strToInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		help(1, "argument must be a number, got:", s)
	}

	return int(i)
}

func readFile(name string) []byte {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalln(err)
	}

	return buf
}

func printAsJson(resp [][]interface{}) {
	buf, err := json.Marshal(resp)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(buf))
}
