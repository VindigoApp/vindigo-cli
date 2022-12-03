ENTRY_NAME=vindigo.go
BINARY_NAME=bin/vindigo

build() {
	LINUX_BIN="${BINARY_NAME}-linux"
	WINDOWS_BIN="${BINARY_NAME}-win.exe"
	DARWIN_BIN="${BINARY_NAME}-darwin"

	echo "Building Linux"
	GOARCH=amd64 GOOS=linux go build -o ${LINUX_BIN} ${ENTRY_NAME}

	echo "Building Windows"
	GOARCH=amd64 GOOS=windows go build -o ${WINDOWS_BIN} ${ENTRY_NAME}

	echo "Building OSX"
	GOARCH=arm64 GOOS=darwin go build -o ${DARWIN_BIN} ${ENTRY_NAME}

	echo "Completing build"
	chmod +x "${LINUX_BIN}"
	chmod +x "${WINDOWS_BIN}"
	chmod +x "${DARWIN_BIN}"
}

clean() {
	echo "Cleaning dist folder"

	go clean
	rm ${BINARY_NAME}*
}

$1;
