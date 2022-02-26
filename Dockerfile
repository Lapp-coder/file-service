FROM golang:1.17.1-alpine3.14 AS builder

COPY ./ /github.com/Lapp-coder/file-service
WORKDIR /github.com/Lapp-coder/file-service

RUN chmod +x wait-for-minio.sh

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -installsuffix "static" -o ./build/bin/file-service ./cmd/main.go

FROM alpine:3.14

WORKDIR /root/

RUN apk --update --no-cache add postgresql-client

COPY --from=builder /github.com/Lapp-coder/file-service/build/bin/file-service .
COPY --from=builder /github.com/Lapp-coder/file-service/wait-for-minio.sh .

CMD ["./file-service"]
