# Builder
FROM golang:1.14 as builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o trigger .

# Runner
FROM alpine:3.11.5

COPY --from=builder /app/trigger /trigger

ENTRYPOINT ["/trigger"]
