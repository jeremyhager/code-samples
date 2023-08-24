GO_CMD=go
GO_BUILD=${GO_CMD} build
BIN_PATH=bin
BINARY_NAME=code-samples

default: build

build: build-linux build-mac

all: clean 
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

build-linux: clean
	${GO_BUILD} GOOS=linux GOARCH=amd64 GOOS=linux --release --target-dir ${BINPATH}/code-samples-linux

build-linux: clean
	${GO_BUILD} GOOS=linux GOARCH=amd64 GOOS=linux --release --target-dir ${BINPATH}/code-samples-mac

clean:
	rm -rf ${BIN_PATH}