FROM golang:alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN apk add build-base
RUN go build -o svc /build/cmd/app/main.go

FROM alpine

COPY --from=builder /build/svc /app/
COPY config/config.json /app/config/config.json
WORKDIR /app

CMD ["./svc"]