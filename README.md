# Hermetic Build Engine

A secure, containerized build runner that enforces network isolation and cryptographic verification of artifacts.

## Architecture
- **Isolation:** Uses Podman with `--network none`.
- **Integrity:** Read-only source mounting and SHA-256 binary hashing.
- **Configurable:** Uses `build.json` for environment and target specifications.

## Quick Start
1. Ensure `Podman` is installed and running.
2. Configure your build in `build.json`.
3. Run the engine:
   ```bash
   go run main.go