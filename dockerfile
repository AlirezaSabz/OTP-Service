# build stage
FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# CGO_ENABLED=0 for static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o otp-service ./main.go

# run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/otp-service .

EXPOSE 8080
CMD ["./otp-service"]
