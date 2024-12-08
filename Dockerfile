# Build stage
FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Run stage
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]
