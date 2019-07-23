FROM        golang:1.12.5
MAINTAINER  yjge
WORKDIR $GOPATH/src/github.com/learn_go
COPY . $GOPATH/src/github.com/learn_go
RUN go build test.go

EXPOSE 8080
ENTRYPOINT ["./test"]