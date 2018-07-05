SHELL:=/bin/bash

test:
	GOPATH=$(shell pwd) go test def/...

build: test
	mkdir -p bin/
	GOPATH=$(shell pwd) go build -o bin/gh_scrap src/main.go
	cp bin/gh_scrap .

run: build
	./bin/gh_scrap

build-docker:
	docker build -t redwhitemiko/gh_scrap .

run-docker: build-docker
	docker run --rm redwhitemiko/gh_scrap .
