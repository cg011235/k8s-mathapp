# Build the Go app
FROM golang:1.16 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mathapp ./cmd/main.go

# Run the Go app
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/mathapp .
EXPOSE 8080
CMD ["./mathapp"]
