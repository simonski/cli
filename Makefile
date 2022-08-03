default_target: build
.PHONY : default_target upload

usage:
	@echo "The goutils Makefile"
	@echo ""
	@echo "Usage : make <command> "
	@echo ""
	@echo "commands"
	@echo ""
	@echo "  clean                 - cleans temp files"
	@echo "  test                  - builds and runs tests"
	@echo "  build                 - creates binary"
	@echo "  run                   - runs in go run"
	@echo "  install               - builds and installs"
	@echo ""
	@echo "  all                   - all of the above"
	@echo ""

clean:
	go clean
	
build: clean test format
	go fmt
	go build
	
test:
	go test -v ./...

install:
	go install

run: 
	go run main.go $*

all: clean build test
	go install

format:
	staticcheck ./...
	go fmt ./...

setup:
	go install honnef.co/go/tools/cmd/staticcheck@latest

