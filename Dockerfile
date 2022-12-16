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

RUN go install github.com/matryer/moq@latest \
    && go install github.com/pressly/goose/v3/cmd/goose@latest \
    && go generate ./... \
    && go env

RUN go build -o /query-monitoring /app/cmd/main.go
RUN chmod +x migrate.sh \
    && mv migrate.sh /opt/migrate.sh

FROM alpine:latest

ENV workspace /app

# Create and switch to a new user
RUN addgroup -S appuser && \
    adduser \
    --disabled-password \
    --gecos "" \
    --home "$(workspace)" \
    --ingroup "appuser" \
    --no-create-home \
    --uid "999" \
    "appuser"

WORKDIR ${workspace}
USER appuser

COPY --from=builder /query-monitoring /app/query-monitoring
COPY --from=builder /opt/migrate.sh /app/migrate.sh
COPY --from=builder /go/bin/goose /app/goose

EXPOSE 8080

ENTRYPOINT [ "/app/migrate.sh" ]

CMD ["/app/query-monitoring"]