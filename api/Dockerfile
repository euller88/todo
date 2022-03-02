FROM golang:1.17-alpine AS builder

RUN apk update && apk add --no-cache ca-certificates git tzdata && update-ca-certificates

WORKDIR /src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go get -d -v

RUN go build -o /app .

EXPOSE 55555

CMD ["/app"]