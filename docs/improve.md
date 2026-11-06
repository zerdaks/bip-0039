# Improve

Leftovers from /improve, so the next run can start here instead of re-auditing. Each line is stamped with the commit it was found at; a run deletes a line once its files have changed or it has acted on it.

- 2026-11-06 | 3b1ef77 | internal/bip39/bip39.go | bytesToBinStr concatenates strings in a loop - `go fix -diff ./...` (Go 1.26+ modernizer) flags it; use a strings.Builder with fmt.Fprintf(&out, "%08b", b).
- 2026-11-06 | 3b1ef77 | internal/bip39/bip39_test.go | TestGenerateMnemonic128Words is misnamed - it checks 128, 192 and 256 bits (12/18/24 words); read the test body.
- 2026-11-06 | 3b1ef77 | go.mod go.sum | testify v1.10.0 is behind v1.12.1 - `go list -m -u all`; `just update` bumps it.
