# Go parameters
GO := go
GOBUILD := $(GO) build
GOCLEAN := $(GO) clean
GOTEST := $(GO) test

# Target binary name
TARGET := dist/image_exif_extractor

.PHONY: all build clean test

all: build

build:
	$(GOBUILD) -o $(TARGET) main.go

clean:
	$(GOCLEAN)
	rm -f $(TARGET)

test:
	$(GO) test -v ./tests/...