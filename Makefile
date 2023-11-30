BIN 					:= gutenberg-ts
VERSION 			:= $(shell git describe --tags)
PREFIX 				:= $(shell go env GOPATH)

.PHONY: build
build: 
	@go build -ldflags="-s -w -X main.AppVersion=${VERSION}" ./cmd/${BIN}

.PHONY: clean
clean: 
	@rm -f ${BIN}
