FROM golang:1.10

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

ADD . $GOPATH/src/samdasu-alddle

WORKDIR $GOPATH/src/samdasu-alddle

RUN go get -u github.com/golang/dep/...
RUN dep ensure

EXPOSE 50051
RUN go build -o server ./samdasu-server
CMD ["./server"]
