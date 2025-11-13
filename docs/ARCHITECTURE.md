# Architecture Notes

## Clean Architecture Layers
- **Domain**: pure value objects with zero dependencies
- **App**: orchestrates domain logic, owns interfaces for persistence/messaging
- **Transport**: converts HTTP concerns into app calls, handles validation + responses
- **Bootstrap**: shared wiring (config, logging, router, graceful shutdown)

## Services
| Service | Purpose | Example Ports |
|---------|---------|---------------|
| Gateway | Aggregates APIs, rate limiting, fan-out | 9000 |
| Auth | JWT/OAuth, roles, identity | 9101 |
| Blog | Authoring, drafts, scheduling | 9102 |
| Interaction | Comments, reactions, reporting | 9103 |
| Notification | Email/WebSocket notifications | 9104 |
| Analytics | Views, engagement metrics, reports | 9105 |
| Bookmark | Reading lists & progress | 9106 |
| Search | Full-text search/indexing | 9107 |
| Worker | Schedules + executes background jobs | 9201 |

## Request Flow Example (Create Comment)
1. Next.js client hits `gateway` → `/blogs/{id}/comments`
2. Gateway validates auth + forwards to Interaction Service
3. Interaction Service `transport/http` decodes payload, calls `app.Service.CreateComment`
4. App service persists (future DB repo) and emits `events.EventCommentCreated`
5. Worker/Notification subscribers react to event for fan-out notifications

## Event Conventions
Event names live in `pkg/events`. Use `events.Envelope` for async payloads with `Name`, `Source`, `Payload`. Suggested routing keys:
- `user.signed_up`
- `blog.published`
- `comment.created`
- `notification.created`
- `analytics.recorded`

## Observability Hooks
- `internal/common/config` exposes Prometheus + tracing endpoints per service
- Extend `pkg/middleware` for request metrics
- Add OTEL exporters under `pkg/observability` (future)

## Deployment Strategy
- Local: `docker compose` builds each service with the shared multi-stage Dockerfile
- CI: run `go test ./...`, then `docker build --build-arg SERVICE=<svc>` per service
- Production: push images to a registry, deploy via Kubernetes/Gateway (NGINX/Envoy) + Observability stack
