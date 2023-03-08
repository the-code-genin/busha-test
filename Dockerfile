FROM golang:1.19.3-alpine

RUN apk add g++

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o build/bin/app .

CMD ["build/bin/app"]