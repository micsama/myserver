# Binary name
BINARY=server
# Builds the project
build:
		go build -o ${BINARY}
		go test -v
# Installs our project: copies binaries
install:
		go install
release:
		# Clean
		go clean
		rm -rf *.gz
		# Build for mac
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
		go build
		# Build for linux
		go clean
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
		# Build for win
		go clean
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
		go clean
# Cleans our projects: deletes binaries
clean:
		go clean

.PHONY:  clean build
