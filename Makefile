default:
	@echo "make install"
	@echo "make test"
	@echo "make lint"
	@echo "make run"
	@echo "make doc"

install: install-deps install-go-deps

install-deps:
	brew install go

install-go-deps:
	go install github.com/mfridman/tparse@latest

# NOTE: -count=1 ignores cached tests
test:
	@echo "Running tests..."
	go test -v -count=1 ./... -json | tparse -all

lint:
	@echo "Formatting code..."
	go fmt ./...

run:
	@echo "Running..."
	go run cmd/main.go

RESET=$(shell printf "\033[0m")
YELLOW=$(shell printf "\033[33m")

doc:
	@echo "Generating documentation..."
	@echo -e "Visit ${YELLOW}http://localhost:6060/pkg/github.com/zerdaks/bip-0039/?m=all${RESET}"
	godoc -http=:6060

.SILENT: default install test lint run doc
