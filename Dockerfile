FROM golang:1.12-alpine3.10

# Copy files from context
WORKDIR /go/src/github.com/pablopb3/biwenger-api/
COPY . .

# Install dep
RUN apk add --no-cache --virtual .build-deps git \
    && go get -u github.com/golang/dep/cmd/dep

# Get dependencies
RUN /go/bin/dep ensure

# Build
RUN go build -o /go/bin/biwenger-api .

# Run
EXPOSE 8080
ENTRYPOINT /go/bin/biwenger-api

