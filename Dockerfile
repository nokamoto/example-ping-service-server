FROM golang:1.11.0-alpine3.8 AS build

RUN apk update && apk add git

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/nokamoto/example-ping-service-server

COPY Gopkg.* ./
COPY *.go ./

RUN dep ensure -vendor-only=true

RUN go install .

FROM alpine:3.8

COPY --from=build /go/bin/example-ping-service-server /usr/local/bin/example-ping-service-server

ENTRYPOINT [ "example-ping-service-server" ]
