# Improve

Leftovers from /improve, so the next run can start here instead of re-auditing. Each line is stamped with the commit it was found at; a run deletes a line once its files have changed or it has acted on it.

- 2026-11-14 | 64cfd62 | internal/bip39/loader.go internal/bip39/bip39.go | Word list entries keep surrounding spaces - LoadWordList does not trim and validate only rejects empty or duplicate words, so `"abandon "` loads and yields a mnemonic with a doubled space; check by loading an fstest.MapFS list with a trailing space on one line.
