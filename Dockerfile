FROM golang:1.18-alpine

WORKDIR /app

ADD . /app

RUN go mod download

RUN go env -w GO111MODULE=on

RUN go build -o /grpc-travel

EXPOSE 5000
EXPOSE 5001


CMD [ "/grpc-travel" ] 