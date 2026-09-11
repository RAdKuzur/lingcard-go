FROM golang:1.27.0-alpine as builder

COPY . .

RUN go build server/main.go

ENTRYPOINT ["./main"]