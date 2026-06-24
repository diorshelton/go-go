# go-go

A from-scratch HTTP/1.1 server in Go using only stdlib. No `net/http` for the core plumbing — raw TCP connections, hand-parsed request bytes, manual routing, and an in-memory cache guarded by an `RWMutex`. Built as a learning project at the boundary of backend and network fundamentals.

---

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

---

## Repo Structure

```
go-go/
├── cmd/server/main.go       # entry point; --port and --workers flags
├── server/
│   ├── listener.go          # TCP accept loop; hands connections to pool
│   ├── pool.go              # fixed-N worker goroutines + keep-alive loop
│   ├── parser.go            # raw HTTP/1.1 parsing (no net/http)
│   ├── router.go            # route registration and matching
│   └── response.go          # HTTP response serialization
├── cache/
│   └── store.go             # RWMutex-protected in-memory KV store
└── handlers/
    └── keys.go              # route handlers for /keys and /keys/:id
```

---

## Getting Started

**Build**
```bash
make build
```

**Run**
```bash
make run
# Starts on port 8080 with 4 workers
```

**Custom flags**
```bash
go run ./cmd/server/main.go --port 9090 --workers 8
```

**Lint** (requires [golangci-lint](https://golangci-lint.run/))
```bash
golangci-lint run
```

---

## Routes

| Method | Path        | Behavior                                                    |
|--------|-------------|-------------------------------------------------------------|
| GET    | /keys       | <!-- TODO: update after #4 is closed (return all key-value pairs vs. key names only) --> |
| GET    | /keys/:id   | Return the value stored at `id`                             |
| PUT    | /keys/:id   | Store a JSON value at `id`                                  |
| DELETE | /keys/:id   | Delete the entry at `id` and return the deleted value       |

All responses include `Content-Type: application/json` and `Content-Length` headers.

> **Note:** Response bodies for error cases (400, 404) are not yet valid JSON — tracked in [#3](https://github.com/diorshelton/go-go/issues/3). This section will be updated with example response shapes once that issue is closed.

---

## Design Constraints

**No external dependencies.** `net` and `sync` from stdlib only. No frameworks, no routers, no middleware libraries.

**No `net/http`.** Connection handling, parsing, and routing are all hand-rolled. That's the point.

**Fixed worker pool.** Concurrency is bounded and explicit. N goroutines are started at startup and pull connections from a channel. Keep-alive loops within each worker — no new goroutines per request.

**Cache uses `json.RawMessage`.** Values are stored and returned as raw bytes, so JSON is never re-encoded on the way out.

**RWMutex disciplines.** Concurrent reads don't block each other (`RLock`/`RUnlock`). Writes are exclusive (`Lock`/`Unlock`).

**Defensive parsing.** Partial reads, malformed input, and unsupported methods return `400 Bad Request` without panicking or leaking goroutines. A 30-second read deadline on each connection handles idle clients.

**Clean shutdown.** `SIGINT`/`SIGTERM` closes the listener, drains in-flight connections, and exits with no goroutine leaks.

---

## Out of Scope

HTTPS/TLS, HTTP/2, pipelining, authentication, persistence, eviction/TTL, file serving, external dependencies.
