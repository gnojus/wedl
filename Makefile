wedl:
	VERSION=$(shell git describe --tag)
	go build -ldflags "-s -w -X main.version=$(shell git describe --tag)"