# Portfolio Backend (Go Microservices)

Learning-focused microservice backend that powers a portfolio/blog experience. It showcases clean architecture, consistent HTTP/JSON transports, shared bootstrap wiring, and a Docker Compose stack you can run as a whole mesh or service-by-service while pairing with a separate Next.js frontend.

## Architecture & Services
- Clean architecture: **domain** (pure structs) → **app** (business orchestration) → **transport/http** (chi handlers) → **module** (registrar) → `/cmd/<service>` (entrypoint) backed by shared **bootstrap/config/server** code.
- Cross-cutting packages: structured logging (`pkg/logger`), HTTP helpers (`pkg/httpx`), request logging middleware (`pkg/middleware`), in-memory cache placeholder (`pkg/cache`), DB placeholder (`pkg/db`), and event naming (`pkg/events`).
- Services (default ports in `internal/common/config`):
  - Gateway `:9000` — aggregates routes and exposes `/health` with discovered upstreams.
  - Auth `:9101` — signup/login/refresh/me.
  - Blog `:9102` — drafts, publish, trending, version history.
  - Interaction `:9103` — comments and reactions.
  - Notification `:9104` — email/websocket-style notifications.
  - Analytics `:9105` — record and fetch blog metrics.
  - Bookmark `:9106` — save/list/delete reading list entries.
  - Search `:9107` — index documents and query search results.
  - Worker `:9201` — enqueues background jobs via HTTP shim.

## Repository Layout
```
.                     # github.com/ajay/portfolio-backend
├── cmd/<service>/    # Entrypoints that call bootstrap.Run(...)
├── internal/
│   ├── common/       # config loader, router + HTTP server bootstrap
│   ├── <service>/    # domain + app + transport + module per service
├── pkg/              # shared helpers (logger, middleware, cache, db, events, httpx)
├── deploy/           # Dockerfile (multi-service) + docker-compose stack
├── docs/             # architecture notes
├── tools/            # install-go helper
└── .env.example      # baseline env vars for local/dev + compose
```

## API Surface (quick reference)
- Gateway: `GET /health`
- Auth: `POST /auth/signup`, `POST /auth/login`, `POST /auth/refresh`, `GET /auth/me`
- Blog: `POST /blogs`, `PUT /blogs/{id}/publish`, `GET /blogs/trending`, `GET /blogs/{id}/history`
- Interaction: `POST /blogs/{id}/comments`, `GET /blogs/{id}/comments`, `POST /blogs/{id}/reactions`
- Notification: `POST /notifications`, `GET /notifications?userId=...`
- Analytics: `POST /analytics`, `GET /analytics/blog/{id}`
- Bookmark: `POST /bookmarks/{blogId}`, `GET /bookmarks?userId=...`, `DELETE /bookmarks/{id}`
- Search: `POST /search/index`, `GET /search?q=term`
- Worker: `POST /tasks/{job}`

Handlers return JSON stubs today; persistence, messaging, and cache layers are intentionally left as extension points.

## Configuration
- Global knobs (`.env.example`): `APP_ENV`, `HTTP_PORT`, `DATABASE_URL`, `CACHE_URL`, `BROKER_URL`, `TRACE_ENDPOINT`, and service URL hints for the gateway (`AUTH_SERVICE_URL`, `BLOG_SERVICE_URL`, etc.).
- Per-service overrides: every service reads env vars with a prefix (e.g., `AUTH_HTTP_PORT`, `BLOG_DATABASE_URL`) via `internal/common/config.Load`. If a prefixed var is absent, global defaults and the service map above are used.
- Logging: `LOG_LEVEL` (`debug|info|warn|error`) drives `pkg/logger` (slog JSON).

## Prerequisites
- Go **1.23+**, GNU `make`
- `docker` and `docker compose` (for the full stack)
- `curl`/`wget` for quick API checks

### Install Go Quickly
```bash
make install-go GO_VERSION=1.23.0
```
The helper downloads Go to `/usr/local/go` (may prompt for sudo) and prints `PATH`/`GOROOT` exports for your shell profile.

## Running Locally (Go tools)
1. Copy envs: `cp .env.example .env` (adjust ports/URLs as needed).
2. Start a service: `make run-auth` (or `run-blog`, `run-gateway`, etc.). The `Makefile` maps `run-<svc>` to `go run ./cmd/<svc>`.
3. Hit an endpoint, e.g.:
   ```bash
   curl -X POST http://localhost:9101/auth/signup \
     -d '{"email":"demo@example.com","password":"pass","name":"Demo"}' \
     -H "Content-Type: application/json"
   ```
4. Format/tests: `make fmt`, `make test` (runs `gofmt` and `go test ./...`).

> Note: Defaults come from `internal/common/config/defaultPorts`; setting `HTTP_PORT` in `.env` forces all services to share that port (compose does this for 8080). Unset it to use per-service ports while running locally.

## Running the Full Mesh (Docker Compose)
1. Build & start from `deploy/`:
   ```bash
   cd deploy
   docker compose up --build gateway
   ```
2. Ports: gateway `9000`, auth `9101`, blog `9102`, interaction `9103`, notification `9104`, analytics `9105`, bookmark `9106`, search `9107`, worker `9201` (all internal containers listen on `8080`).
3. Health check:
   ```bash
   curl http://localhost:9000/health
   ```
   returns upstream service URLs registered in the gateway.

## How the Pieces Fit Together
- Bootstrap (`internal/common/bootstrap`): loads config, constructs a slog logger, sets up a chi router with shared middleware, registers service routes, and starts an HTTP server with graceful shutdown handling for SIGINT/SIGTERM.
- Config (`internal/common/config`): builds `ServiceConfig` structs with HTTP/DB/cache/messaging/observability sections, driven by env vars and sensible defaults.
- Server (`internal/common/server`): chi router + standard library `http.Server` with timeouts and basic request logging; easy to extend with more middleware in `pkg/middleware`.
- Domain/App layers: each service exposes small domain structs (e.g., `auth/domain.Credentials`, `blog/domain.Post`) and app services that encapsulate the business flows (ID generation, timestamps, log statements).
- Transports: HTTP handlers decode/validate JSON, call app services, and write responses via `pkg/httpx`.
- Events: `pkg/events` lists canonical event names (`user.signed_up`, `blog.published`, etc.) and the `Envelope` shape to keep async messaging consistent when real brokers are added.

## Walkthrough: Analytics Service (end-to-end)
Use this service as the template for reading the others—the structure is identical.

1. **Entrypoint** (`cmd/analytics/main.go`): imports only `bootstrap` and the analytics `module`. `main()` calls `bootstrap.Run("analytics", module.Registrar())`; bootstrap handles config/log/router/server while the module wires service-specific parts.
2. **Module/Registrar** (`internal/analytics/module/module.go`): exposes a `Registrar` func (matching the type in `internal/common/bootstrap`). It receives the resolved `config.ServiceConfig`, a scoped logger, and the chi router. Inside it constructs the app service (`app.NewService(log)`) and passes it to the HTTP transport's `Register` helper.
3. **Domain layer** (`internal/analytics/domain/metrics.go`): holds plain structs shared between transports + business logic. JSON tags (`json:"blogId"`) ensure request/response bodies map cleanly onto the Go struct.
4. **App layer** (`internal/analytics/app/service.go`): implements methods like `Record` and `Fetch`. The file only imports `context`, `log/slog`, `time`, and the domain package, which keeps business logic free from HTTP concerns. `NewService` is where real adapters (DB/cache) would be injected; comments in the file call this out.
5. **Transport layer** (`internal/analytics/transport/http/handler.go`): imports chi for routing, encoding/json for decoding payloads, and the shared `pkg/httpx` helper for consistent responses. `Register` hangs two routes off `/analytics`. `handler.record` shows the typical flow—decode JSON, call app service, write response/error—while `handler.fetch` demonstrates path params with `chi.URLParam`.

Because every service follows the same layout you can swap "analytics" with "auth", "blog", etc., and navigate confidently: `cmd` → `module` → `app`/`domain`/`transport`. The inline comments recently added to those files mirror the explanation above.

## Extending This Scaffold
- Wire real adapters: replace `pkg/db.NopStore` and `pkg/cache.MemoryCache` with PostgreSQL/Redis clients; inject them into app services.
- Add messaging: integrate NATS/RabbitMQ/Kafka using `pkg/events`, push events from app services, and subscribe in the worker/notification services.
- Observability: expose Prometheus metrics and OTEL traces using the hooks in `config.ObservabilityConfig`.
- Harden transports: add auth middleware, validation, rate limiting (gateway is the natural home), and richer error contracts.

## Troubleshooting
- Ports already in use: set `HTTP_PORT` (or `<SERVICE>_HTTP_PORT`) to a free port before running.
- Env not loading: ensure `.env` is present or variables are exported in your shell. `config.Load` pulls from the environment; there is no `.env` parsing baked in.
- Builds slow in Docker: take advantage of the Go build cache mounts already present in `deploy/Dockerfile`; avoid `--no-cache` unless necessary.

Happy hacking! Feel free to plug in real storage/messaging and evolve the stubs as you go.
