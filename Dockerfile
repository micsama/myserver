FORM golang:latest
MAINTAINER Dzmfg "csgomicsma@gmail.com"
WORKDIR $GOPATH/src/github.com/micsama/myserver
ADD . /Users/dzmfg/WorkSpace/myserver
RUN go build 
EXPOSE 6064
ENTRYPOINT ["./server"]
