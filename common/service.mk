#################################################################################
#								variables										#
#################################################################################
# Basic go commands
GOCMD=go
GOGET=$(GOCMD) get
GOCLEAN=$(GOCMD) clean
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build

# code coverage variables
COVER_MODE=count # put metrics about how many times did each statement run.
COVER_PROFILE=coverage.txt # coverage profile, write out file.
COVER_REPORT=coverage.html # view the coverage report(.html file) in the browser.

# Build variables
BUILD_FILE=main.go
BUILD_BINARY=main

# Deployment bundle variables

ZIP_FILE=main.zip
ZIP_FILE_LIST=$(BUILD_BINARY) 

# Outliner
GREEN=$(shell tput -Txterm setaf 2)
YELLOW=$(shell tput -Txterm setaf 3)
RESET=$(shell tput -Txterm sgr0)
TARGET_MAX_CHAR_NUM=20

#################################################################################
#									rules										#
#################################################################################
.DEFAULT_GOAL := help

help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = $$1; sub(/:$$/, "", helpCommand); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

######################### environment target rules #################################
## Run rules for the development sign-off
dev: clean cover build mod prepare

############################# targets split up ######################################
## Remove temporary files
clean:
	@echo "\n*************** cleaning up ***************\n"
	rm -f $(BUILD_BINARY)
	rm -f $(ZIP_FILE)

## Run go mod tidy & verify dependencies
mod:
	@echo "\n*************** removing unused dependencies ***************\n"
	$(GOMOD) tidy
	
	@echo "\n*************** verifying dependencies ***************\n"
	$(GOMOD) verify

# -mod=readonly: Prohibits go build from modifying go.mod
## Run all the tests
test:
	@echo "\n*************** running unit test(s) ***************\n"
	$(GOTEST) -mod=$(MOD) -covermode=$(COVER_MODE) -coverprofile=$(COVER_PROFILE) ./...

## Run all the tests & prepare a coverage report in .html file
cover: test
	@echo "\n*************** preparing test coverage report ***************\n"
	# $(GOGET) -u github.com/axw/gocov/gocov
	# $(GOGET) -u github.com/matm/gocov-html
	gocov convert $(COVER_PROFILE) | gocov-html > $(COVER_REPORT)

# GOOS=linux GOARCH=amd64: Preparing a binary to deploy to AWS Lambda requires that it is compiled for Linux
# -mod=readonly: Prohibits go build from modifying go.mod
# -s: Omit the symbol table and debug information.
# -compressdwarf: Compress DWARF if possible (default true).
## Build a version
build:
	@echo "\n*************** building ***************\n"
	rm -f $(BUILD_BINARY)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_BINARY) -mod=$(MOD) -ldflags="-s -compressdwarf" $(BUILD_FILE)

## Prepare the build for the deployment
prepare:
	@echo "\n*************** packing binary & dependency file(s) for the deployment ***************\n"
	rm -f $(ZIP_FILE)
	zip -j $(ZIP_FILE) $(ZIP_FILE_LIST)
