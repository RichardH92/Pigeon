FROM golang

EXPOSE 8081

ENV GOPATH=/go

RUN cd $GOPATH/src
RUN git clone https://github.com/RichardH92/Pigeon.git
RUN cd Pigeon
RUN go get -t -v ./...
RUN go install ./main/...

ENTRYPOINT [$GOPATH/bin/main]

