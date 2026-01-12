FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o minesweeper main.go
FROM alpine:latest
WORKDIR /game
COPY --from=builder /app/minesweeper .
ENV TERM=xterm-256color
ENTRYPOINT ["./minesweeper"]