# Halo Runtime

A Deno-like JavaScript runtime built in Go, powered by QuickJS.

## Project Structure

```
edon/
├── cmd/                    # Application entry points
│   ├── edon/               # Main CLI (REPL, file execution, package management)
│   │   ├── main.go
│   │   ├── init.go
│   │   └── npm.go
│   ├── runtime/            # Standalone runtime CLI
│   │   └── main.go
│   └── web/                # Web-based REPL server
│       ├── main.go
│       └── static/
├── internal/               # Private application code
│   ├── modules/
│   │   ├── console/        # Console API implementation
│   │   └── loader/         # Module loading, NPM, resolution
│   ├── runtime/            # Core JS runtime
│   └── server/             # HTTP server for web REPL
├── tests/
│   ├── integration/
│   ├── unit/
│   └── fixtures/
├── Dockerfile
├── makefile
└── go.mod
```

## Getting Started

### Build

```bash
make build
make build-web
make build-all
```

### Run

```bash
./bin/halo
./bin/halo script.js
./bin/halo -eval "console.log('Hello!')"
./bin/halo init
./bin/halo install lodash
```

### Development

```bash
go run ./cmd/edon

go run ./cmd/web

make test
```

## Features

- **REPL** - Interactive JavaScript shell with history and autocomplete
- **File Execution** - Run `.js` files directly
- **Web REPL** - Browser-based JavaScript playground
- **NPM Support** - Install and use NPM packages
- **Module Loading** - Support for local, CDN, and NPM imports

## Roadmap

- [ ] Module caching system
- [ ] URL import parsing
- [ ] Module resolution for URL imports
- [ ] JSR registry support
