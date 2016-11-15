FROM golang

EXPOSE 8081

ENV GOPATH=/go

CMD go get github.com/gorilla/mux
CMD go get github.com/aws/aws-sdk-go/aws
CMD go get github.com/jmespath/go-jmespath

ENTRYPOINT ["/go/setup.sh"]

