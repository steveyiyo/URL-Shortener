FROM golang:1.18-alpine3.15 as builder

RUN apk --update upgrade \
    && apk add --no-cache sqlite gcc g++ \
    && rm -rf /var/cache/apk/*

WORKDIR /app
COPY . .
RUN  CGO_ENABLED=1 CC=gcc go build -o /app/app main.go

FROM alpine:3.10

# Install libc6-compat for go-sqlite3 use
RUN apk --update upgrade \
    && apk add --no-cache sqlite \
    && rm -rf /var/cache/apk/*

COPY --from=builder /app/app /app/
WORKDIR /app
CMD ["./app"]