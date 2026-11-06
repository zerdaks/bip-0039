# Improve

Leftovers from /improve, so the next run can start here instead of re-auditing. Each line is stamped with the commit it was found at; a run deletes a line once its files have changed or it has acted on it.

- 2026-11-06 | 3876b06 | internal/bip39/bip39.go | binToInt error branch is unreachable - its input is always 11 chars from bytesToBinStr, so the ParseInt error path in MnemonicFromEntropy can never run (shows as uncovered). Check: `go test -coverprofile` and read the binToInt call in MnemonicFromEntropy.
