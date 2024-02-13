.PHONY: build
build:
	go build -v ./cmd

run:
	make build
	.\cmd.exe

.DEFAULT_GOAL := build