SHELL := /bin/bash

default:
	@echo "make test"
	@echo "make lint"
	@echo "make run"
	@echo "make doc"

define color_pass
$(shell printf "\033[32mPASS\033[0m")
endef

define color_fail
$(shell printf "\033[31mFAIL\033[0m")
endef

# Run tests and colorize the output
test:
	@echo "Running tests..."
	go test -v ./... 2>&1 | \
		sed 's/PASS/$(color_pass)/g' | \
		sed 's/FAIL/$(color_fail)/g'

lint:
	@echo "Formatting code..."
	go fmt ./...

run:
	@echo "Running..."
	go run .

doc:
	@echo "Generating documentation..."
	godoc -http=:6060

.PHONY: default test lint run doc
.SILENT: default test lint run doc