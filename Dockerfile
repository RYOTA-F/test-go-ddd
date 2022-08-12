FROM golang

RUN mkdir /go/src/work

WORKDIR /go/src/work

ADD . /go/src/work

CMD [ "go", "run", "api/main.go" ]
