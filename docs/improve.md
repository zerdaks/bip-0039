# Improve

Leftovers from /improve, so the next run can start here instead of re-auditing. Each line is stamped with the commit it was found at; a run deletes a line once its files have changed or it has acted on it.

- 2026-11-06 | 5e7bb53 | internal/bip39/bip39.go internal/bip39/bip39_test.go | checksum takes a redundant entLen - `checksum(entropy, entLen)` could derive the length from `len(entropy)*8`; the tests rely on the mismatch (`checksum([]byte{0x04}, 128)`) and so test inputs no caller can produce. Check: grep for `checksum(` call sites.
- 2026-11-06 | 5e7bb53 | internal/bip39/bip39.go | binToInt error branch is unreachable - its input is always 11 chars from bytesToBinStr, so the ParseInt error path in MnemonicFromEntropy can never run (shows as uncovered). Check: `go test -coverprofile` and read bip39.go:44-47.
