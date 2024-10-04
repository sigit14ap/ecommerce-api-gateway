FROM golang:1.21.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /api-gateway ./cmd/main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /api-gateway .

COPY .env .

#EXPOSE 7000
# Cloud Run expects the app to listen on port 8080
EXPOSE 8080

CMD ["./api-gateway"]
