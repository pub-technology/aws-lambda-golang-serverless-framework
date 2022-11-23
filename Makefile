.PHONY: build clean deploy

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/product/get handlers/product/get/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/product/getAll handlers/product/getAll/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/product/delete handlers/product/delete/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/product/put handlers/product/put/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

format:
	gofmt -w handlers/product/get/main.go
	gofmt -w handlers/product/getAll/main.go
	gofmt -w handlers/product/delete/main.go
	gofmt -w handlers/product/put/main.go
