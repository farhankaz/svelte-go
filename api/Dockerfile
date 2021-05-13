# syntax=docker/dockerfile:1
FROM golang:1.16 as builder
WORKDIR /go/src/github.com/farhankaz/svelte-go
COPY main.go .
RUN go GOARCH=386 GOOS=linux go build -o app main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/farhankaz/svelte-go/app .
CMD ["./app"]  