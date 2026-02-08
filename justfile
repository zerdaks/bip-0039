# Install dependencies
[group('setup')]
install:
    @brew install go
    @go install golang.org/x/tools/cmd/goimports@latest
    @go install github.com/mfridman/tparse@latest
    @go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Update dependencies
[group('setup')]
update:
    @echo "Updating dependencies..."
    @go get -u ./...
    @go mod tidy

# Run tests
[group('dev')]
test:
    @echo "Running tests..."
    @go test -v -count=1 ./... -json | tparse -all -format plain # -count=1 disables test caching

# Run tests for package
[group('dev')]
testpkg dir="":
    @echo "Running tests in directory {{dir}}..."
    @go test -v -count=1 ./{{dir}}/...

# Format and vet code
[group('dev')]
lint:
    @echo "Formatting and vetting code..."
    @go fmt ./...
    @go vet ./...
    @golangci-lint run ./...

# Run cmd applications
[group('dev')]
run:
    @echo "Running..."
    @go run cmd/main.go

RESET := `printf "\033[0m"`
YELLOW := `printf "\033[33m"`

# Run documentation server
[group('dev')]
doc:
    @echo "Generating documentation..."
    @echo "{{YELLOW}}http://localhost:6060/pkg/github.com/zerdaks/bip-0039/?m=all{{RESET}}"
    @godoc -http=:6060
