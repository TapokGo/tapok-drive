ARG GO_VERSION=1.25
ARG ALPINE_VERSION=3.22

# Stage 1: Build Go binary
FROM golang:${GO_VERSION}-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./bin/tapokdrive ./cmd/tapokdrive

# Stage 2: Create a runtime image
FROM alpine:${ALPINE_VERSION}

RUN adduser -D -s /bin/sh appuser

COPY --from=builder /app/bin/tapokdrive /usr/local/bin/tapokdrive

USER appuser

EXPOSE 8080

CMD [ "/usr/local/bin/tapokdrive" ]