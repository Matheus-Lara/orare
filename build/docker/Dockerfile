FROM golang:1.23-alpine AS builder

ARG APP_MODE
ENV GIN_MODE=${APP_MODE:-"release"}
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./orare ./cmd/orare/main.go

FROM alpine:latest AS runner
ARG APP_MODE
ENV GIN_MODE=${APP_MODE:-"release"}
WORKDIR /app
COPY --from=builder /app/orare .

ENTRYPOINT ["./orare"]