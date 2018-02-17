FROM golang:latest as builder
WORKDIR /go/src/github.com/beandrop/joiner/
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/beandrop/joiner/app .
CMD ["./app"]