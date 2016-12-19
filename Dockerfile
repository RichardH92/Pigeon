FROM golang

EXPOSE 8081

ENV GOPATH=/go
ENV GOBIN=/go/bin

RUN cd $GOPATH/src; git clone https://github.com/RichardH92/Pigeon.git
WORKDIR $GOPATH/src/Pigeon/pigeon
RUN go get -t -v ./...
RUN go install ./...
RUN cp ../config.json ../../../bin/config.json

CMD $GOPATH/bin/pigeon

