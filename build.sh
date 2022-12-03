ENTRY_NAME=vindigo.go
BINARY_NAME=bin/vindigo

build() {
	LINUX_BIN="${BINARY_NAME}-linux"
	WINDOWS_BIN="${BINARY_NAME}-win.exe"
	DARWIN_BIN="${BINARY_NAME}-darwin"

	F_PREFIX="github.com/VindigoApp/vindigo-cli/utils"
	F_VERSION="0.1.0"
	F_COMMIT=$(git rev-parse --short HEAD)
	F_DATE=$(date -u '+%Y-%m-%d_%I:%M:%S%p')

	FLAGS="-X ${F_PREFIX}.Version=$F_VERSION -X ${F_PREFIX}.Commit=$F_COMMIT -X ${F_PREFIX}.Date=$F_DATE -X ${F_PREFIX}.Production=true"

	echo "Building Linux"
	GOARCH=amd64 GOOS=linux go build -ldflags "$FLAGS" -o ${LINUX_BIN} ${ENTRY_NAME}

	echo "Building Windows"
	GOARCH=amd64 GOOS=windows go build -ldflags "$FLAGS" -o ${WINDOWS_BIN} ${ENTRY_NAME}

	echo "Building OSX"
	GOARCH=arm64 GOOS=darwin go build -ldflags "$FLAGS" -o ${DARWIN_BIN} ${ENTRY_NAME}

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