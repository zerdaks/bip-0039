BOLD := `printf "\033[1m"`
RED := `printf "\033[31m"`
YELLOW := `printf "\033[33m"`
RESET := `printf "\033[0m"`

default:
    @just --list

# Install dependencies
[group('setup')]
install:
    @brew install go
    @go install golang.org/x/tools/cmd/goimports@latest
    @go install golang.org/x/tools/cmd/godoc@latest
    @go install github.com/mfridman/tparse@latest
    @curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.10.1

# Update dependencies
[group('setup')]
update:
    @echo "{{RED}}{{BOLD}}Updating dependencies...{{RESET}}"
    @go get -u ./...
    @go mod tidy

# Run tests
[group('dev')]
test:
    @echo "{{RED}}{{BOLD}}Running tests...{{RESET}}"
    @go test -v -count=1 ./... -json | tparse -follow-verbose -all -format plain # -count=1 disables test caching

# Run tests for package
[group('dev')]
testpkg dir="":
    @echo "{{RED}}{{BOLD}}Running tests in directory {{dir}}...{{RESET}}"
    @go test -v -count=1 ./{{dir}}/...

# Format and vet code
[group('dev')]
lint:
    @echo "{{RED}}{{BOLD}}Formatting and vetting code...{{RESET}}"
    @go fmt ./...
    @go vet ./...
    @golangci-lint run ./...

# Run cmd applications
[group('dev')]
run:
    @echo "{{RED}}{{BOLD}}Running...{{RESET}}"
    @go run cmd/main.go

# Run documentation server
[group('dev')]
doc:
    @echo "{{RED}}{{BOLD}}Generating documentation...{{RESET}}"
    @echo "{{YELLOW}}http://localhost:6060/pkg/github.com/zerdaks/bip-0039/?m=all{{RESET}}"
    @godoc -http=:6060
