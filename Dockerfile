FROM golang:latest

WORKDIR /go/src/github.com/pablopb3/biwenger-api/

# Install dep
RUN apk add --no-cache --virtual .build-deps git \
    && go get -u github.com/golang/dep/cmd/dep

RUN /go/bin/dep ensure

ADD . .
RUN go build -o /go/bin/biwenger-api .

EXPOSE 8080
ENTRYPOINT /go/bin/biwenger-api

