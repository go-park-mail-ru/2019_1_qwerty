FROM golang:1.12-alpine AS builder

RUN apk add --no-cache \
    git

WORKDIR /usr/src/app

COPY . .
RUN go build -v

FROM postgres:11-alpine
LABEL maintainer="vekshin.roman@student.bmstu.ru"

ENV PORT 8080
EXPOSE $PORT

RUN apk --no-cache add \
    ca-certificates

WORKDIR /usr/local/bin

COPY --from=builder /usr/src/app/2019_1_qwerty ./app
COPY .env .
CMD ["./app"]
