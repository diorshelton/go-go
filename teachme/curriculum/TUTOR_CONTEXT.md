# Tutor Context — go-go Curriculum

This file provides context for resuming tutoring sessions. Read it at the start of any session where the learner continues this curriculum.

## Learner Profile

- **Name:** Dior
- **Level:** Junior Go engineer — has built CRUD apps in Go, comfortable with the language, new to low-level networking
- **Goal:** Portfolio-quality understanding of HTTP internals + interview-readiness on concurrency, TCP, and server architecture
- **Learning style:** Hands-on, Socratic — prefers to think before being given answers. Asks "why" questions frequently. Responds well to code examples with one specific thing to reason about.
- **Tone:** Serious, focused. No fluff. Explain the tradeoff, not just the rule.
- **Schedule:** ~2 hours/day, 4-week target

## Curriculum Structure

4 modules, ~16 lessons, ~13 exercises. All files are in `curriculum/` as plain HTML, opened via `file://`.

| Module | Topic | Status |
|--------|-------|--------|
| 01 | Network Foundation (TCP, sockets, accept loop) | Complete |
| 02 | HTTP as Raw Bytes (request line, headers, body, keep-alive) | Complete |
| 03 | Concurrency in Go (goroutines, channels, worker pool, RWMutex, race conditions) | Complete |
| 04 | Building go-go (listener, parser, router, cache, response, edge cases) | Complete |

## Progress Tracking

Progress is tracked via `localStorage` in each HTML file. Keys follow the pattern `gogo_m{module}_{type}{number}` — e.g., `gogo_m01_l01` for a lesson, `gogo_m03_e02` for an exercise. The dashboard reads these keys to compute completion percentages.

## What go-go is

A from-scratch HTTP/1.1 server in Go using only stdlib (`net`, `sync`). No `net/http`, no external dependencies. The intended file layout:

```
cmd/server/main.go       -- --port, --workers flags
server/
  listener.go            -- TCP accept loop, graceful shutdown
  pool.go                -- fixed-N worker goroutines + keep-alive loop
  parser.go              -- raw HTTP/1.1 parsing
  router.go              -- route matching and handlers
  response.go            -- HTTP response formatting
cache/
  store.go               -- RWMutex-protected in-memory KV store
```

## Key Concepts Covered

- **TCP:** bidirectional byte stream, 3-way handshake, OS backlog queue, `net.Listen`/`Accept()`
- **HTTP/1.1:** request line format, header parsing with CRLF terminator, body framing via Content-Length, keep-alive default behavior, `bufio.Reader` reuse
- **Concurrency:** goroutine lifecycle, buffered channels, fixed worker pool with `sync.WaitGroup`, RWMutex (RLock for reads, Lock for writes), data race definition, race detector (`go test -race`)
- **go-go implementation:** each component's responsibility, the edge cases that break naive implementations, clean shutdown sequence

## Tutoring Notes

- Dior responds well to being asked "why" before being told the answer — lead with the Socratic question
- Don't over-explain. State the principle, give one concrete example, stop.
- If stuck, ask: "what do you think happens if you don't do X?" rather than explaining X directly
- The goal is interview-readiness — if explaining a concept, frame it as "here's how you'd explain this to an interviewer"
- Don't introduce abstractions beyond what go-go needs. The point is explicitness, not elegance.
