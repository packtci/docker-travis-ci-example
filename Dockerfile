FROM golang:1.10-alpine

RUN mkdir /go/src/app

WORKDIR /go/src/app

COPY . .

RUN go build -o main .

CMD ["/go/src/app/main"]