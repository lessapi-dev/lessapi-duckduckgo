# Stage 1: Modules caching
FROM golang:1.21 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Stage 2: Build
FROM golang:1.21 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /workdir
WORKDIR /workdir
RUN GOOS=linux GOARCH=amd64 go build -o /bin/lessapi-duckduckgo ./cmd/main

# Stage 3: Final
FROM lessapi/base-env-ubuntu-playwright-go:ubuntu22.04-pwgo-v0.4201.1
COPY --from=builder /bin/lessapi-duckduckgo /
COPY ./resource /resource
RUN chmod +x /lessapi-duckduckgo
CMD ["/lessapi-duckduckgo"]