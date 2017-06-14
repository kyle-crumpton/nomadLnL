FROM golang:1.8

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get && \
    go build

EXPOSE 8080

ENTRYPOINT ["/go/bin/app"]

CMD []
