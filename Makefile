.PHONY: build
build:
	go build -v ./cmd/schedule
	
.DEFAULT_GOAL := build