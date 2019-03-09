FROM golang:1.9-alpine3.7 AS builder

RUN apk add --no-cache \
    git

WORKDIR /usr/src/app

COPY . .
RUN go get -d
RUN go build -v

FROM postgres:11-alpine
LABEL maintainer="vekshin.roman@student.bmstu.ru"

EXPOSE 8080

RUN apk --no-cache add \
    ca-certificates

WORKDIR /usr/local/bin

COPY --from=builder /usr/src/app/app .
COPY .env .
CMD ["./app"]
