FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /app

RUN rm -rf ~/.netrc
RUN go build -o /docker-day-7-revision

FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder docker-day-7-revision .
COPY .env /app

EXPOSE 8080
CMD ["./docker-day-7-revision"]