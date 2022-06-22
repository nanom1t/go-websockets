FROM golang:1.18-alpine

MAINTAINER nanom1t

ENV PORT 3000

WORKDIR /go/src/app
COPY . /go/src/app

RUN apk --no-cache add git
RUN go mod download && go build -o app

EXPOSE $PORT

ENTRYPOINT ["./app"]
