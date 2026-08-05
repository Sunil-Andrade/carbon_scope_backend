# ---------- Build Stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# ---------- Runtime Stage ----------
FROM alpine:3.22

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

COPY .env .

EXPOSE 8080

CMD ["./main"]