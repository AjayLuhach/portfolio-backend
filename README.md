# Portfolio Backend (Go Microservices)

A clean-architecture Go backend showcasing a multi-service blogging platform for portfolio and learning use cases. The repo focuses on backend concerns (frontend is delivered by a separate Next.js project).

## Features at a Glance
- Auth, blog, interaction, notification, analytics, bookmark, search, gateway, and worker services.
- Shared bootstrap layer for configuration, structured logging, HTTP server, and graceful shutdown.
- Event-naming conventions, lightweight in-memory cache + DB placeholders, and HTTP helper utilities.
- Docker Compose stack for local orchestration plus a generic multi-service Dockerfile.

## Prerequisites
- `curl` or `wget`
- `docker` + `docker compose`
- GNU `make`

## Install Go Globally
Use the provided helper to install Go 1.23 system-wide (falls back to sudo if required):

```bash
make install-go GO_VERSION=1.23.0
```

This script downloads Go from `go.dev`, installs it under `/usr/local/go`, and prints the PATH exports to drop into your shell profile. Update `GOROOT`, `GOPATH`, and `PATH` as instructed after it runs.

## Project Layout
```
.
├── cmd/                 # Service entrypoints (auth, blog, gateway, ...)
├── internal/
│   ├── common/          # Shared config, bootstrap, HTTP server wiring
│   ├── auth|blog|...    # Service specific domain/app/transport layers
├── pkg/                 # Cross-cutting helpers (logger, middleware, cache)
├── deploy/              # Dockerfile + docker-compose for the full stack
├── tools/               # Tooling scripts (Go installer)
├── .env.example         # Baseline environment variables
└── README.md
```

Each service follows the same pattern:
- `domain`: pure structs/value objects
- `app`: business logic orchestrators
- `transport/http`: HTTP handlers powered by Chi
- `module`: exposes a bootstrap registrar consumed by `/cmd/<service>`

## Running Locally
1. Copy the env template: `cp .env.example .env`
2. Start individual services: `make run-auth` (or `run-blog`, `run-gateway`, etc.)
3. Orchestrate everything via Docker Compose:
   ```bash
   cd deploy
   docker compose up --build gateway
   ```
   Gateway listens on `http://localhost:9000`, downstream services on `9101-9107` and the worker on `9201`.

Sample endpoints (JSON stubs for now):
- `POST /auth/signup`
- `POST /blogs`
- `GET /blogs/trending`
- `POST /blogs/{id}/comments`
- `GET /notifications`
- `GET /analytics/blog/{id}`
- `POST /bookmarks/{blogId}`
- `GET /search?q=term`
- `POST /tasks/reindex`

## Next Steps
- Connect PostgreSQL/Redis/Search adapters inside each service (`pkg/db`, `pkg/cache` act as placeholders).
- Add background job scheduling (e.g., Asynq, NATS, or RabbitMQ) and emit real events defined in `pkg/events`.
- Expand tests with `testing`, `testify`, and `httptest` packages per service.
- Wire observability (Prometheus metrics, OpenTelemetry tracing, Loki logging) using the config hooks already in place.

Happy hacking!
