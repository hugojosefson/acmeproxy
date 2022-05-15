FROM golang:1 as builder
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build the actual app
COPY . .
RUN go build -v -o dist/acmeproxy .

##############################################################################
FROM debian:stable-slim as runner
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/dist/acmeproxy /usr/bin/acmeproxy
RUN chmod +x /usr/bin/acmeproxy
ENTRYPOINT [ "/usr/bin/acmeproxy" ]
