# Build stage
FROM golang:1.22-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-data-api-microservices-new ./...

# Runtime stage
FROM alpine:3.19

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/go-data-api-microservices-new /app/go-data-api-microservices-new

EXPOSE 8080

CMD ["/app/go-data-api-microservices-new"]
