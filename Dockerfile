FROM golang:1.23-alpine AS builder
WORKDIR /go/src
ENV GO111MODULE=on

COPY ./go.mod ./
RUN go mod download

COPY ./main.go .

RUN CGO_ENABLED=0 GOOS=linux go build cgo -o web-main

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /go/src/web-main .

EXPOSE 8080
ENTRYPOINT ["./web-main"]