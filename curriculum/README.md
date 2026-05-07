# go-go Curriculum

A personalized learning curriculum for building a from-scratch HTTP/1.1 server in Go.

## How to use this

Open `dashboard.html` in your browser. Everything runs as `file://` — no server, no build step, no accounts.

```
open curriculum/dashboard.html
```

Work through the modules in order. Each lesson has "Mark complete" buttons that track your progress in localStorage. The dashboard shows your progress across all modules.

## Structure

```
curriculum/
├── dashboard.html                        # Your home base — open this first
├── curriculum.json                       # Machine-readable curriculum state
├── TUTOR_CONTEXT.md                      # Context for Claude tutor sessions
│
├── module-01-network-foundation/
│   ├── lesson-01-tcp-connections.html
│   ├── lesson-02-three-way-handshake.html
│   ├── lesson-03-net-listen-accept.html
│   └── exercises/
│       ├── exercise-01-tcp-quiz.html
│       └── exercise-02-tcp-echo-server.html
│
├── module-02-http-raw-bytes/
│   ├── lesson-01-request-line.html
│   ├── lesson-02-headers.html
│   ├── lesson-03-body-content-length.html
│   ├── lesson-04-keep-alive.html
│   └── exercises/
│       ├── exercise-01-request-dissector.html
│       ├── exercise-02-spot-malformed.html
│       └── exercise-03-http-quiz.html
│
├── module-03-concurrency/
│   ├── lesson-01-goroutines-channels.html
│   ├── lesson-02-worker-pools.html
│   ├── lesson-03-rwmutex.html
│   ├── lesson-04-race-conditions.html
│   └── exercises/
│       ├── exercise-01-pool-design.html
│       └── exercise-02-rwmutex-drill.html
│
└── module-04-building-go-go/
    ├── lesson-01-listener.html
    ├── lesson-02-parser.html
    ├── lesson-03-router.html
    ├── lesson-04-cache.html
    ├── lesson-05-response.html
    ├── lesson-06-edge-cases.html
    └── exercises/
        ├── exercise-01-listener.html
        ├── exercise-02-parser.html
        ├── exercise-03-router.html
        ├── exercise-04-cache.html
        ├── exercise-05-response.html
        └── exercise-06-wire-together.html
```

## What you're building

go-go is a from-scratch HTTP/1.1 server using only Go stdlib. The intended layout once you complete the exercises:

```
go-go/
├── cmd/server/main.go
├── server/
│   ├── listener.go
│   ├── pool.go
│   ├── parser.go
│   ├── router.go
│   └── response.go
└── cache/
    └── store.go
```

Run it: `go run ./cmd/server/main.go --port 8080 --workers 10`

Test it: `go test -race ./...`
