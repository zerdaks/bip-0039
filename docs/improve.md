# Improve

Leftovers from /improve, so the next run can start here instead of re-auditing. Each line is stamped with the commit it was found at; a run deletes a line once its files have changed or it has acted on it.

- 2026-11-05 | 1458f6d | justfile | Go toolchain unpinned against golangci-lint pin - `install` runs `brew install go` (latest) while golangci-lint is pinned, so a future Go release will break `just lint` again exactly as Go 1.27 broke v2.10.1; check by installing the pinned golangci-lint and running it against the latest brew Go.
