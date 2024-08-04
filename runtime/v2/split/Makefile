CLI_DIR = cmd

SHIMV2 = containerd-shim-split-v2
SHIMV2_OUTPUT = $(CURDIR)/$(SHIMV2)
SHIMV2_DIR = $(CLI_DIR)
TARGET_DIR = /usr/local/bin

# LDFLAGS for build customization
KATA_LDFLAGS ?= -s -w
BUILDFLAGS ?=

# Find all Go source files
SOURCES := $(shell find . -name '*.go')

# Build the shim
.PHONY: all build install clean

all: $(SHIMV2_OUTPUT)

$(SHIMV2_OUTPUT): $(SOURCES) 
	@echo "Building $(SHIMV2)..."
	cd $(SHIMV2_DIR) && go build -ldflags "$(KATA_LDFLAGS)" $(BUILDFLAGS) -o $(SHIMV2_OUTPUT)

install: $(SHIMV2_OUTPUT)
	@echo "Installing $(SHIMV2) to $(TARGET_DIR)..."
	@install -m 0755 $(SHIMV2_OUTPUT) $(TARGET_DIR)

clean:
	@echo "Cleaning up..."
	@rm -f $(SHIMV2_OUTPUT)

