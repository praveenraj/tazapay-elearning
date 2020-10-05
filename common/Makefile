#################################################################################
#								variables										#
#################################################################################
# Basic go commands
GOCMD=go
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOCOV=gocov

# code coverage variables
COVER_MODE=count # put metrics about how many times did each statement run.
COVER_PROFILE=coverage.txt # coverage profile, write out file.
COVER_REPORT=coverage.html # view the coverage report(.html file) in the browser.

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

## Prepare the library package
library: cover mod

## Run all the tests
test:
	@echo "\n*************** running unit test(s) ***************\n"
	$(GOTEST) -covermode=$(COVER_MODE) -coverprofile=$(COVER_PROFILE) ./...

## Run all the tests & prepare a coverage report in .html file
cover: test
	@echo "\n*************** preparing test coverage report ***************\n"
	$(GOGET) -u github.com/axw/gocov/gocov
	$(GOGET) -u github.com/matm/gocov-html
	$(GOCOV) convert $(COVER_PROFILE) | gocov-html > $(COVER_REPORT)

## Run go mod tidy & verify dependencies
mod:
	@echo "\n*************** removing unused dependencies ***************\n"
	$(GOMOD) tidy
	
	@echo "\n*************** verifying dependencies ***************\n"
	$(GOMOD) verify

## Build all services (ex: make build-all-svc ENV=[dev,stage,release,prod]). If ENV value is not provided then it will exe with the default value
build-all-svc:
	@for dir in $(shell find ../ -name "*service" -type d); do \
		echo '\n${YELLOW}##############################' $${dir} '##############################${RESET}\n'; \
		cd $${dir}; \
		make $(ENV); \
	done
