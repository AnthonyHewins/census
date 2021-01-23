.PHONY: clean $(targets)

targets := cli

$(targets):
	go build -o bin/$@ cmd/$@/main.go

clean:
	find *.go -type f -exec gofmt -w -s {} \;
	go mod tidy
	rm -rf ./bin

get-vars:
	@curl https://api.census.gov/data/$(year)/$(service)/$(version)/variables.json | grep concept | sort | uniq -u
