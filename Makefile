
BUILD_OUTPUT_FILENAME ?= butta

# platform-specific settings
GOOS := $(shell go env GOOS)

ifeq ($(GOOS),windows)
	ifneq ($(suffix $(BUILD_OUTPUT_FILENAME)),.exe)
		BUILD_OUTPUT_FILENAME := $(addsuffix .exe,$(BUILD_OUTPUT_FILENAME))
	endif
endif

.PHONY: build
build:
	go build -o dist/api/$(BUILD_OUTPUT_FILENAME) $(BUILD_FLAGS) ./cmd/api
	go build -o dist/console/$(BUILD_OUTPUT_FILENAME) $(BUILD_FLAGS) ./cmd/console

.PHONY: dev
dev:
	test -f .env && . .env; air -c .air.toml .

.PHONY: install
install:
	@# install Air verse
	go install 'github.com/air-verse/air@latest'
