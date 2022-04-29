FROM golang:alpine

COPY . /app
WORKDIR /app

EXPOSE 8080

RUN go build -o main

CMD ["./main"]
