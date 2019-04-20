FROM golang:1.12-alpine AS builder

RUN apk add --no-cache \
    build-base \
    git

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build

FROM alpine
LABEL maintainer="vekshin.roman@student.bmstu.ru"

EXPOSE $PORT

RUN apk --no-cache add \
    ca-certificates

WORKDIR /usr/local/bin

COPY sql/ sql/
COPY --from=builder /usr/src/app/2019_1_qwerty ./app
CMD ["./app"]
