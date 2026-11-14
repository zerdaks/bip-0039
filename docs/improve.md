# Improve

Leftovers from /improve, so the next run can start here instead of re-auditing. Each line is stamped with the commit it was found at; a run deletes a line once its files have changed or it has acted on it.

- 2026-11-14 | f0a60cf | internal/bip39/loader.go internal/bip39/loader_test.go | LoadWordList keeps trailing whitespace - a CRLF word list loads as 2048 words ending in `\r` and passes validate; check by loading an fstest.MapFS file with `\r\n` line endings.
