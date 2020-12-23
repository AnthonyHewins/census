package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/url"
	"github.com/AnthonyHewins/census"
)

func genericQuery(path string, params string) {
	form := url.Values{}

	if params != "" {
		var mapObj map[string]interface{}

		if err := json.Unmarshal([]byte(params), &mapObj); err != nil {
			log.Fatalln(err)
		}

		for key, val := range mapObj {
			form.Add(key, fmt.Sprint(val))
		}
	}

	resp, err := census.Query(path, form)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(resp))
}