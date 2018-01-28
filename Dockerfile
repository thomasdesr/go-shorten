FROM golang:1.9

EXPOSE 8080

ARG package=github.com/thomaso-mirodin/go-shorten
# ARG package=app 
# ${PWD#$GOPATH/src/}

RUN mkdir -p /go/src/${package}
WORKDIR /go/src/${package}

COPY . /go/src/${package}
RUN go-wrapper download
RUN go-wrapper install

RUN go get "github.com/GeertJohan/go.rice/rice"
RUN go generate

CMD ["go-wrapper", "run"]
