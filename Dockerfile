FROM golang:latest
ADD . /go/src/app

 WORKDIR /go/src/app
 ENV GOPATH /go
 ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
 RUN go get "gopkg.in/mgo.v2/bson"
 RUN go get "gopkg.in/mgo.v2"
 RUN go get "github.com/tidwall/gjson"
 RUN go get "github.com/gorilla/mux"
 RUN go build -o /go/bin/biwenger-api .
 EXPOSE 8080

 ENTRYPOINT /go/bin/biwenger-api

