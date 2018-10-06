FROM golang:latest

MAINTAINER harunwols@gmail.com (harun)

WORKDIR /go/src/ntm

COPY . .

RUN go get -u github.com/kardianos/govendor
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega

RUN govendor sync

