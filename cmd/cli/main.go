package main

import (
	"os"
	"fmt"
	"encoding/json"
	"net/url"
	"log"

	"github.com/AnthonyHewins/census"
)

func main() {
	n := len(os.Args)

	switch n {
	case 2:
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help" {
			help(0)
		}

		buildQuery(os.Args[1], "")
	case 3:
		buildQuery(os.Args[1], os.Args[2])
	default:
		help(1, "wrong number of args:", n)
	}
}

func buildQuery(path string, params string) {
	form := url.Values{}

	if params != "" {
		var mapObj map[string]interface{}
		err := json.Unmarshal([]byte(params), &mapObj)
		if err != nil {
			log.Fatalln(err)
		}

		for key, val := range mapObj {
			form.Add(key, fmt.Sprint(val))
		}
	}

	fmt.Println(census.Query(path, form))
}

const helpText = `usage: census COMMAND
where COMMAND expands to one of:

  -h, --help, help	Print this help message

  PATH [JSON]		Make a request to PATH and use the string JSON
					as the URL-encoded form for the request. Leaving
					JSON blank is equivalent to having no form
`

func help(exitCode int, extraMessages ...interface{}) {
	fmt.Println(extraMessages...)
	fmt.Println(helpText)
	os.Exit(exitCode)
}
