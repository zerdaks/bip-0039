BOLD := `printf "\033[1m"`
RED := `printf "\033[31m"`
RESET := `printf "\033[0m"`

default:
    @just --list

# Install dependencies
[group('setup')]
install: githooks
    @brew install go
    @brew install betterleaks
    @go install github.com/mfridman/tparse@latest
    @curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.15.0

# Install Git hooks
[group('setup')]
githooks:
    @git config core.hooksPath hooks

# Update dependencies
[group('setup')]
update:
    @echo "{{RED}}{{BOLD}}Updating dependencies...{{RESET}}"
    @go get -u -t ./... # -t includes test-only dependencies such as testify
    @go mod tidy

# Run tests
[group('dev')]
test:
    @echo "{{RED}}{{BOLD}}Running tests...{{RESET}}"
    @go test -count=1 -cover ./... -json | tparse -progress -format plain # -count=1 disables test caching

# Run tests for package
[group('dev')]
testpkg dir="":
    @echo "{{RED}}{{BOLD}}Running tests in directory {{dir}}...{{RESET}}"
    @go test -v -count=1 -cover ./{{dir}}/...

# Lint and check formatting
[group('dev')]
lint:
    @echo "{{RED}}{{BOLD}}Linting code...{{RESET}}"
    @# golangci-lint runs go vet among its defaults.
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
    @go doc -http
