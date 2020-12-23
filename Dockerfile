FROM golang:1.14-buster

WORKDIR /sauce
ADD selenium_simple.go /sauce
ENV GOBIN=$GOPATH/bin
RUN go get ./...
RUN go build -o ./selenium_simple ./selenium_simple.go

CMD ["/sauce/selenium_simple"]
