default:
	@echo "make test"
	@echo "make lint"
	@echo "make run"
	@echo "make doc"

GREEN=$(shell printf "\033[32m")
RED=$(shell printf "\033[31m")
RESET=$(shell printf "\033[0m")
YELLOW=$(shell printf "\033[33m")

test:
	@echo "Running tests..."
	go test -v ./... 2>&1 | \
		sed "s/PASS/$(GREEN)PASS$(RESET)/g" | \
		sed "s/FAIL/$(RED)FAIL$(RESET)/g"

lint:
	@echo "Formatting code..."
	go fmt ./...

run:
	@echo "Running..."
	go run cmd/main.go

doc:
	@echo "Generating documentation..."
	@echo -e "Visit ${YELLOW}http://localhost:6060/pkg/github.com/zerdaks/bip-0039/?m=all${RESET}"
	godoc -http=:6060

.PHONY: default test lint run doc
.SILENT: default test lint run doc
