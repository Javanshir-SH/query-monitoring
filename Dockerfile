FROM golang:1.19.0-alpine3.15 as builder

RUN apk update\
 && apk add --no-cache make\
 && apk add --no-cache ca-certificates\
 && update-ca-certificates

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /query-monitoring /app/cmd/main.go

FROM alpine:latest
COPY --from=builder /query-monitoring /query-monitoring
EXPOSE 8080
ENTRYPOINT ["/query-monitoring"]