.PHONY: build clean

# Version information
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "0.1.0")
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Go build flags
LDFLAGS = -ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.GitCommit=${GIT_COMMIT}"

# Build the application
build:
	@echo "Building go-arc version ${VERSION}"
	@go build ${LDFLAGS} -o go-arc

# Install the application
install: build
	@echo "Installing go-arc"
	@mv go-arc ${GOPATH}/bin/go-arc

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f go-arc

# Run tests
test:
	@go test ./...

# Create a new release tag
release:
	@echo "Current version is ${VERSION}"
	@read -p "Enter new version (e.g., v0.2.0): " version; \
	git tag -a $$version -m "Release $$version"
	@echo "Run 'git push origin --tags' to push the new tag"
