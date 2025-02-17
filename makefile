# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)


# ==============================================================================
# Chat

run: 
	go run api/services/cap/main.go | go run api/tooling/logfmt/main.go

test-status:
	curl -i -X GET http://localhost:3000/status


# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor