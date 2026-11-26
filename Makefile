PREFIX ?= $(HOME)/go/bin

.PHONY: install test

install:
	mkdir -p "$(PREFIX)"
	go build -o "$(PREFIX)/ox" ./cmd

test:
	go test ./...
