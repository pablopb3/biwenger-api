FROM golang:latest

WORKDIR /go/src/github.com/pablopb3/biwenger-api/
RUN go get "gopkg.in/mgo.v2/bson" && \
	go get "gopkg.in/mgo.v2" && \
	go get "github.com/tidwall/gjson" && \
	go get "github.com/gorilla/mux" && \
	go get "github.com/magiconair/properties" && \
	go get "github.com/dghubble/go-twitter/twitter" && \
	go get "github.com/dghubble/oauth1"

ADD . .
RUN go build -o /go/bin/biwenger-api .

EXPOSE 8080
ENTRYPOINT /go/bin/biwenger-api

