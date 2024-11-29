FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /app/main

FROM scratch
COPY --from=builder /app/main /app/main
ENTRYPOINT ["/app/main"]