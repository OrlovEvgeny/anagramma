BINARY := anagramma

.PHONY: linux
linux:
	mkdir -p build/linux
	GOOS=linux GOARCH=amd64 go build -o build/linux/$(BINARY) cmd/main.go

.PHONY: darwin
darwin:
	mkdir -p build/osx
	GOOS=darwin GOARCH=amd64 go build -o build/osx/$(BINARY) cmd/main.go

.PHONY: build
build:  linux darwin