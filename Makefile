.PHONY: clean $(targets)

targets := cli

$(targets):
	go build -o bin/$@ cmd/$@/main.go

clean:
	find . -iname *.go -type f -exec gofmt -w -s {} \;
	go mod tidy
	rm -r ./bin
