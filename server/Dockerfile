FROM golang:1.14.7-buster

ENV GO111MODULE=on

EXPOSE 8080

WORKDIR /go/src/app

CMD ["./app"]

COPY . .

RUN go vet && go build -o app
