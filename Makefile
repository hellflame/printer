
BINARY_PATH = ./dist
BINARY_NAME = printer

tidy:
	find . -name "*.go" -type f -not -path "./vendor/*" | xargs -n1 go fmt
	go mod tidy

build: tidy clean build-darwin build-linux build-windows

dist-path:
	mkdir -p ${BINARY_PATH}

build-darwin: dist-path
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ${BINARY_PATH}/${BINARY_NAME}
	cd ${BINARY_PATH} && zip ${BINARY_NAME}-darwin.zip ${BINARY_NAME} && rm ${BINARY_NAME}

build-linux: dist-path
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ${BINARY_PATH}/${BINARY_NAME}
	cd ${BINARY_PATH} && zip ${BINARY_NAME}-linux.zip ${BINARY_NAME} && rm ${BINARY_NAME}

build-windows: dist-path
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ${BINARY_PATH}/${BINARY_NAME}.exe

clean:
	rm -rf ${BINARY_PATH}
