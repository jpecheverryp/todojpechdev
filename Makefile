## audit: run quality control checks
.PHONY: audit
audit: 
	go fmt ./...
	go mod tidy
	go mod verify
	go vet ./...

## dev: run application in developent mode
.PHONY: dev
dev:
	go run ./cmd/web/

## build: build production binary
.PHONY: build
build:
	npm run build
	GOOS='linux' GOARCH='amd64' go build -o ./bin/web ./cmd/web/
