FROM golang:1.21 AS builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goexpert-stress-test main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/goexpert-stress-test .
ENTRYPOINT ["./goexpert-stress-test"]