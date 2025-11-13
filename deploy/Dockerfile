# syntax=docker/dockerfile:1.6

ARG GO_VERSION=1.23
FROM golang:${GO_VERSION} AS base
WORKDIR /src
COPY go.mod go.sum* ./
RUN --mount=type=cache,target=/go/pkg/mod \
    if [ -f go.sum ]; then go mod download; fi
COPY . .

ARG SERVICE=auth
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /out/service ./cmd/${SERVICE}

FROM gcr.io/distroless/base-debian12 AS runtime
WORKDIR /app
COPY --from=base /out/service /app/service
EXPOSE 8080
ENTRYPOINT ["/app/service"]
