FROM golang:latest

WORKDIR /go/src/app
RUN go get "gopkg.in/mgo.v2/bson" && \
	go get "gopkg.in/mgo.v2" && \
	go get "github.com/tidwall/gjson" && \
	go get "github.com/gorilla/mux"
ADD . .
RUN go build -o /go/bin/biwenger-api .

EXPOSE 8080
ENTRYPOINT /go/bin/biwenger-api

