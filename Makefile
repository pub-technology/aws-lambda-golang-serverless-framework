.PHONY: build clean deploy

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/test src/test.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list src/list.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

format:
	gofmt -w todos/test.go
	gofmt -w todos/list.go
