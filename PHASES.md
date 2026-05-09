# Phases

## Phase 1 — Foundation
Set up the module, directory structure, and entry point. Write the cache (`cache/store.go`) and wire `cmd/server/main.go` so the binary builds and accepts `--port` and `--workers` flags. No server logic yet — just a binary that compiles and starts.

**Done when:** `go build ./...` succeeds and `go run ./cmd/server/main.go --port 8080 --workers 4` starts without panicking.

---

## Phase 2 — Connection Layer
Implement the TCP listener (`listener.go`) and fixed worker pool (`pool.go`). The listener accepts connections and hands them to the pool. Workers pull connections from a channel. Clean shutdown on SIGINT/SIGTERM.

**Done when:** The server accepts a raw TCP connection (e.g. via `nc localhost 8080`), does not crash, and shuts down cleanly on Ctrl-C with no goroutine leaks.

---

## Phase 3 — HTTP Parsing and Routing
Write the HTTP parser (`parser.go`) and router (`router.go`). The parser reads raw bytes off the connection and produces a structured request. The router matches method + path to a handler. Malformed input returns `400 Bad Request`.

**Done when:** A valid `GET /keys` request sent via `curl` returns a `200` response body (even if it's just `{}`), and a malformed request returns `400` without crashing the server.

---

## Phase 4 — Handlers, Keep-Alive, and Edge Cases
Implement the four key routes in `handlers/keys.go` and the response writer (`response.go`). Add keep-alive support — the worker loops on a connection until `Connection: close` or a read timeout. Handle unexpected disconnects without goroutine leaks.

**Done when:** All four routes work correctly end-to-end (`GET /keys`, `GET /keys/:id`, `PUT /keys/:id`, `DELETE /keys/:id`), keep-alive connections persist across multiple requests, and a timed-out or abruptly-closed connection is handled cleanly.
