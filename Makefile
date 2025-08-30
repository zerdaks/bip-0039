default:
	@echo "make install"
	@echo "make update"
	@echo "make test"
	@echo "make lint"
	@echo "make run"
	@echo "make doc"

install: install-deps install-go-deps

install-deps:
	brew install go

install-go-deps:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/mfridman/tparse@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

update:
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

# NOTE: -count=1 ignores cached tests
test:
	@echo "Running tests..."
	go test -v -count=1 ./... -json | tparse -all -format plain

lint:
	@echo "Formatting and vetting code..."
	go fmt ./...
	go vet ./...
	golangci-lint run ./...

run:
	@echo "Running..."
	go run cmd/main.go

RESET=$(shell printf "\033[0m")
YELLOW=$(shell printf "\033[33m")

doc:
	@echo "Generating documentation..."
	@echo -e "Visit ${YELLOW}http://localhost:6060/pkg/github.com/zerdaks/bip-0039/?m=all${RESET}"
	godoc -http=:6060

.SILENT: default install update test lint run doc
