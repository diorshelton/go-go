# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**go-go** is a from-scratch HTTP/1.1 server in Go using only stdlib (`net`, `sync`). No `net/http` for the core plumbing — raw TCP connections, hand-parsed request bytes, manual routing, and an in-memory cache guarded by an RWMutex. It's a learning project; explicitness and observability are intentional design values over abstraction.

## Commands

```bash
# Run the server
make run

# Build
make build

# Lint (requires golangci-lint)
golangci-lint run
```

## Architecture

```
TCP listener (net.Listen)
        ↓
Worker pool — N fixed goroutines + keep-alive loop per connection
        ↓
HTTP parser — method, path, headers, body (defensive, hand-rolled)
        ↓
Router — method + path → handler ←→ Cache (RWMutex KV store)
        ↓
Response writer
```

**Intended layout:**

```
httpserver/
├── cmd/server/main.go       # flags: --port, --workers
├── server/
│   ├── listener.go          # TCP accept loop; hands connections to pool
│   ├── pool.go              # fixed-N worker goroutines + keep-alive loop
│   ├── parser.go            # raw HTTP/1.1 parsing (no net/http)
│   ├── router.go            # route registration and matching
│   └── response.go          # HTTP response formatting
├── cache/
│   └── store.go             # RWMutex-protected in-memory KV store
└── handlers/
    └── keys.go              # GET/PUT/DELETE /keys/:id, GET /keys
```

## Key Design Constraints

- **No external dependencies.** stdlib only — `net` and `sync` are the only relevant packages for the core server.
- **No `net/http`** for connection handling, parsing, or routing. That's the point.
- **Fixed worker pool** — concurrency is bounded and explicit. Workers pull connections from a channel; keep-alive loops within a worker, not by spawning new goroutines per request.
- **Cache uses `json.RawMessage`** as the value type so values are stored and returned without re-encoding.
- **RWMutex disciplines:** concurrent reads don't block each other; writes are exclusive. `RLock`/`RUnlock` for GET, `Lock`/`Unlock` for PUT/DELETE.
- **Defensive parsing:** partial reads, malformed input, and read timeouts must return `400 Bad Request` without panicking or leaking goroutines.
- **Clean shutdown on interrupt** — listener closes, workers drain, no goroutine leaks.

## Routes

| Method | Path          | Behavior                        |
|--------|---------------|---------------------------------|
| GET    | /keys         | Return all keys                 |
| GET    | /keys/:id     | Return value for key            |
| PUT    | /keys/:id     | Set value for key (JSON body)   |
| DELETE | /keys/:id     | Delete key                      |

All responses are JSON with correct `Content-Length` and `Content-Type: application/json` headers.

## Explicitly Out of Scope

HTTPS/TLS, HTTP/2, pipelining, authentication, persistence, eviction/TTL, file serving, external deps.
