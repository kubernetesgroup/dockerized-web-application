FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server main.go

FROM alpine:3.19
RUN addgroup -g 10001 -S appgroup && adduser -u 10001 -S appuser -G appgroup
WORKDIR /home/appuser
COPY --from=builder --chown=appuser:appgroup /app/server .
USER 10001
EXPOSE 5000
ENTRYPOINT ["./server"]
