# Build Stage
FROM golang:1.12

ENV GO111MODULE=on

RUN mkdir -p /go/src/github.com/Molsbee/blog
COPY go.mod /go/src/github.com/Molsbee/blog
COPY go.sum /go/src/github.com/Molsbee/blog
RUN go mod download

COPY . /go/src/github.com/Molsbee/blog

RUN go install github.com/Molsbee/blog

EXPOSE 8080
ENTRYPOINT /go/bin/blog