FROM golang:1.8

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get && \
    go build

ENTRYPOINT ["/go/bin/app"]

CMD []