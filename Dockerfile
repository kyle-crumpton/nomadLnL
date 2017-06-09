FROM golang:1.8

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get 

ENTRYPOINT ["go"]

CMD ["run", "main.go"]