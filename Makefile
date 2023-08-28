GO_CMD=go
GO_BUILD=$(GO_CMD) build
BIN_PATH=bin
BINARY_NAME=code-samples
GOARCH_x64=amd64

default: test build

build: build-linux build-mac build-windows

build-linux: clean
	GOARCH=$(GOARCH_x64) GOOS=linux $(GO_BUILD) -o $(BIN_PATH)/code-samples-linux

build-mac: clean
	GOARCH=$(GOARCH_x64) GOOS=darwin $(GO_BUILD) -o $(BIN_PATH)/code-samples-mac

build-windows: clean
	GOARCH=$(GOARCH_x64) GOOS=windows $(GO_BUILD) -o $(BIN_PATH)/code-samples-windows.exe

clean:
	rm -rf $(BIN_PATH)

test:
	$(GO_CMD) test ./...

fmt:
	$(GO_CMD) fmt ./...