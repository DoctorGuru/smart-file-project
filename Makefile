APP_NAME = organizer
MAIN_PATH = ./cmd/organizer

# Default target
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build     - Build binary for current OS"
	@echo "  make run       - Run the program locally"
	@echo "  make install   - Install globally"
	@echo "  make clean     - Remove build artifacts"
	@echo "  make release   - Build binaries for Linux, Windows, Mac"

# Build binary for current OS
.PHONY: build
build:
	go build -o $(APP_NAME) $(MAIN_PATH)

# Run without building
.PHONY: run
run:
	go run $(MAIN_PATH) -dir="./Downloads" -mode=type

# Install globally (adds to $GOPATH/bin)
.PHONY: install
install:
	go install $(MAIN_PATH)

# Clean build files
.PHONY: clean
clean:
	rm -f $(APP_NAME) $(APP_NAME).exe

# Cross-platform release build
.PHONY: release
release:
	GOOS=linux   GOARCH=amd64 go build -o $(APP_NAME)-linux $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 go build -o $(APP_NAME).exe $(MAIN_PATH)
	GOOS=darwin  GOARCH=arm64 go build -o $(APP_NAME)-mac $(MAIN_PATH)
