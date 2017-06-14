FROM golang:1.8

ENV VERSION "0.0.1"

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get && \
    go build -ldflags "-X main.version=0.0.1" 

EXPOSE 8080

ENTRYPOINT ["/go/bin/app"]

CMD []
