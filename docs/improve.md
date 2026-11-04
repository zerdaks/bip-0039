# Improve

Leftovers from /improve, so the next run can start here instead of re-auditing. Each line is stamped with the commit it was found at; a run deletes a line once its files have changed or it has acted on it.

- 2026-11-03 | 5e4b13e | README.md | BIP-0039 link points at the word list - the link text names the spec but the URL is bip-0039/english.txt; the spec is bip-0039.mediawiki. Check by opening the link.
- 2026-11-03 | 5e4b13e | justfile | goimports installed but unused - `install` runs go install goimports, yet no recipe calls it and lint uses go fmt; check with `grep goimports justfile`. May be an editor dependency, so ask before removing.
