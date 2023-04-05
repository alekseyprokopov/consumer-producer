.PHONY: build
build:
	go build cmd/main.go

.PHONY: test
test:
	go test -v ./...
