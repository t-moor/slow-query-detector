FROM golang:1.18.2-alpine3.15 as builder

RUN apk update\
 && apk add --no-cache make\
 && apk add --no-cache ca-certificates\
 && update-ca-certificates

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /slow-query-detector /app/cmd/main.go

FROM alpine:latest
COPY --from=builder /slow-query-detector /slow-query-detector
EXPOSE 8080
ENTRYPOINT ["/slow-query-detector"]