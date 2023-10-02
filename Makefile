.PHONY: build clean deploy

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/ffcron main.go

clean:
	rm -rf ./bin