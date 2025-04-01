BINARY_NAME=cbn
BINARY_DIR=bin/cbn
CBN_DIR=cmd/cbn

# change default shell (for using source)
SHELL := /bin/bash

install:
	# download all server packages
	go mod download

.PHONY: build
build:
	# build for current platform
	go build -o ${BINARY_DIR}/${BINARY_NAME} main.go

.PHONY: build_all
build_all:
	echo "Building binaries for Windows, Linux macOS (intel), macOS (silicon)"
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_DIR}/${BINARY_NAME}-darwin-silicon main.go
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_DIR}/${BINARY_NAME}-darwin-intel main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_DIR}/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_DIR}/${BINARY_NAME}-windows main.go

# run - just execute the server binary file
.PHONY: run
run:
	./${BINARY_DIR}/${BINARY_NAME}

# build and run server (with some useful cleanup/format commands to always
# keep code clean)
.PHONY: build_and_run
build_and_run: build run

