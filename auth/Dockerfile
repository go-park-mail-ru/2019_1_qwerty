# FROM golang:1.12-alpine AS builder

# RUN apk add --no-cache \
#     build-base \
#     git

# WORKDIR /usr/src/app

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# COPY . .
# RUN go build -o app

FROM ubuntu
LABEL maintainer="vekshin.roman@student.bmstu.ru"

EXPOSE $PORT

RUN apt update && apt install -y \
    ca-certificates

WORKDIR /usr/local/bin

# COPY --from=builder /usr/src/app/app ./app
COPY auth .
CMD ["./auth"]
