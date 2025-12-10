FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY cmd internal index.html go.mod .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

FROM docker.io/debian:bookworm-slim

#RUN apt-get update && apt-get install -y \
#    ca-certificates \
#    tzdata \
#    && rm -rf /var/lib/apt/lists/*

#ENV TZ=Europe/Moscow

WORKDIR /app
COPY --from=builder /app/ocserv-admin .
COPY --from=builder /app/index.html .

EXPOSE 80

#HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 CMD curl -f http://localhost/ || exit 1

CMD ["./ocserv-admin"]
