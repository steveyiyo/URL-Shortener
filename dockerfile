FROM golang as builder
RUN     mkdir -p /app
WORKDIR /app
COPY . .
RUN  go build -o app

FROM scratch
COPY --from=builder /app/app /app
CMD ["/app"]