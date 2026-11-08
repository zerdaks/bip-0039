# Improve

Leftovers from /improve, so the next run can start here instead of re-auditing. Each line is stamped with the commit it was found at; a run deletes a line once its files have changed or it has acted on it.
- 2026-11-08 | 59c6232 | internal/bip39/bip39.go | Unreachable binToInt error path - MnemonicFromEntropy parses 11-char chunks of a string built only from %08b output, so strconv.ParseInt can never fail; the error branch (the only uncovered lines in bip39.go per go tool cover -func) could go, e.g. by computing the index with bit shifts.
