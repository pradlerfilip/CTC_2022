FROM golang:1.16

WORKDIR /go/src/app
COPY ./src .

RUN go get -d -v
RUN go build -v

CMD ["./cv04"]
