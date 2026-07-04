# ==========================
# Build stage
# ==========================
FROM golang:1.25-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build \
  -tags netgo \
  -ldflags="-s -w" \
  -o server \
  ./cmd/app/

# ==========================
# Runtime stage
# ==========================
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /

COPY --from=builder /app/server .

ENV GIN_MODE=release

EXPOSE 8000

ENTRYPOINT ["/server"]
