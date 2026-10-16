# spotify_to_tidal shell

Standalone Go scaffold for a future cross-platform desktop shell and music-provider service contracts. This directory is intentionally a shell only: stubs and interfaces, not a working sync product.

It sits beside the existing Python application and does not replace or depend on it.

## Scope

- Buildable command entrypoint under `cmd/shell`
- Desktop OS stubs for macOS (`darwin`), Linux, and Windows via Go build tags
- Canonical domain models and composable service interfaces (auth, playlist read/write, track match)
- Public provider interface packages for Spotify, TIDAL, and YouTube Music (metadata + contract only)

## Layout

```
shell/
  cmd/shell/           # command entrypoint
  platform/            # Shell interface + OS-specific stubs
  services/domain/     # canonical models and shared contracts
  services/spotify/    # Spotify provider surface
  services/tidal/      # TIDAL provider surface
  services/youtubemusic/ # YouTube Music provider surface
```

## Requirements

- Go 1.22+
- No third-party module dependencies

## Build / run / test

From this directory:

```bash
go build -buildvcs=false -o bin/shell ./cmd/shell
go run -buildvcs=false ./cmd/shell
go test ./...
```

Cross-compile platform stubs:

```bash
GOOS=darwin  go build -buildvcs=false -o bin/shell-darwin  ./cmd/shell
GOOS=linux   go build -buildvcs=false -o bin/shell-linux   ./cmd/shell
GOOS=windows go build -buildvcs=false -o bin/shell-windows.exe ./cmd/shell
```

`-buildvcs=false` avoids VCS stamping when this tree sits inside another git repo.

On the host OS, `go run ./cmd/shell` prints the platform id and exits with a not-implemented error from the stub `Run` method.

## Non-goals

This scaffold does **not** include:

- OAuth or any authentication flow implementation
- HTTP clients or provider API calls
- Track-matching algorithms
- Sync engine or playlist transfer logic
- GUI toolkit or native windowing
- Persistence, secrets, or credential storage
- Generated code
- Code copied from SyncDisBoi, Portify, or other reference projects
- Changes to the existing Python application or its docs
