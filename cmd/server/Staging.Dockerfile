FROM alpine as builder

RUN apk add -U --no-cache ca-certificates

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /app
COPY . /app 

ENTRYPOINT ["/app/server", "--http-port", "8082", "--grpc-port",  "22002", "--project-id", "weather-alert-staging"]
