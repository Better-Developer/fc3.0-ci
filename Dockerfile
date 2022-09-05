FROM golang:latest

WORKDIR /usr/src/app

COPY ./goapp .

RUN go build sum.go

CMD [ "./sum" ]

