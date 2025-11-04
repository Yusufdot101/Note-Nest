# build stage
FROM golang:1.25.3-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main ./api

# final stage
FROM alpine:latest

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/main .

EXPOSE 8080

CMD ["./main"]
