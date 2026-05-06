# HTTP/1.1 Server from Scratch — Go-Go

A from-scratch HTTP/1.1 server in Go with no `net/http` for the core plumbing. Accepts raw TCP connections, parses request bytes by hand, routes to handlers, and serves an in-memory cache backed by a read/write mutex. Built with a fixed worker pool so the concurrency model is explicit and observable. Designed as a learning project at the boundary of backend and network fundamentals.

---

## Learning Objectives

- Understand how a TCP connection is established and how a server accepts and manages multiple connections at the socket level
- Understand the 3-way handshake conceptually — how it works and what Go's `net` package handles on your behalf
- Understand what an HTTP/1.1 request looks like as raw bytes and how to parse it by hand
- Understand HTTP/1.1 connection persistence — how keep-alive works, when connections close, how to manage connection state in the worker loop
- Build a mental model of how routing works without a framework abstracting it away
- Understand concurrent data access — why it's a problem and how RWMutex solves it
- Understand how to handle malformed and unexpected input gracefully — defensive parsing, read timeouts, unexpected disconnects
- Practice explicit concurrency in Go — worker pools, goroutines, channels as plumbing you control

---

## Architecture

```
TCP listener (net.Listen)
        ↓
Worker pool — N fixed goroutines + keep-alive loop
        ↓
HTTP parser — method, path, headers, body (defensive)
        ↓
Router — match path to handler ←→ Cache (RWMutex KV store)
        ↓
Response writer
```

---

## Components

| Component | What it does | Hard part |
|---|---|---|
| TCP listener | `net.Listen`, accept connections, hand off to pool | Lifecycle management, clean shutdown |
| Worker pool | Fixed N goroutines pulling from a connection channel | Bounded concurrency, keep-alive loop per connection |
| HTTP parser | Read raw bytes, parse request line, headers, body | Partial reads, malformed input, read timeouts |
| Router | Match method + path to handler | Clean API design — this becomes the library seam |
| Cache | In-memory map with RWMutex, values as `json.RawMessage` | Concurrent reads vs exclusive writes |
| Response writer | Format and write valid HTTP/1.1 responses | Status lines, headers, content-length correctness |

---

## Key Libraries

`net` and `sync` from stdlib only. No external dependencies. That's the point.

---

## Repo Structure

```
httpserver/
├── cmd/server/main.go       -- flags: --port, --workers
├── server/
│   ├── listener.go          -- TCP accept loop
│   ├── pool.go              -- worker pool + keep-alive loop
│   ├── parser.go            -- raw HTTP parsing
│   ├── router.go            -- route registration + matching
│   └── response.go          -- response formatting
├── cache/
│   └── store.go             -- RWMutex KV store
└── handlers/
    └── keys.go              -- GET, PUT, DELETE /keys/:id
```

---

## Done When

- Server accepts and handles multiple concurrent connections via a fixed worker pool
- HTTP/1.1 requests are parsed from raw bytes without `net/http`
- Four routes work correctly: `GET /keys/:id`, `PUT /keys/:id`, `DELETE /keys/:id`, `GET /keys`
- Keep-alive connections are handled — worker loops on the same connection, closes on `Connection: close` or read timeout
- Malformed requests return `400 Bad Request` without crashing
- Unexpected client disconnects are handled cleanly — no goroutine leaks
- Concurrent reads don't block each other, writes are exclusive
- Responses are valid JSON with correct HTTP status codes
- Server shuts down cleanly on interrupt

---

## Explicitly Out of Scope

- HTTPS / TLS
- HTTP/2
- HTTP/1.1 pipelining
- UDP
- Authentication or authorization
- Persistence — cache lives in memory only, dies on restart
- Eviction policies, TTL, cache expiry
- File I/O / static asset serving
- Production error handling and logging
- External dependencies — stdlib only
