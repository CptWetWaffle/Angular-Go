FROM golang:1.15-alpine AS build
WORKDIR /src/SimpleApp/api

ENV GO111MODULE=on
ENV PORT=3000
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o /bin

EXPOSE 3000 3000
ENTRYPOINT ["/bin/api"]
