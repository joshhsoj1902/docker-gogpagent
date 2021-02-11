FROM golang:latest@sha256:9fdb74150f8d8b07ee4b65a4f00ca007e5ede5481fa06e9fd33710890a624331
RUN mkdir -p /go/src/github.com/joshhsoj1902/docker-gogpagent
ADD . /go/src/github.com/joshhsoj1902/docker-gogpagent/
WORKDIR /go/src/github.com/joshhsoj1902/docker-gogpagent
RUN cp /go/src/github.com/joshhsoj1902/docker-gogpagent/overrides/server.go /go/src/github.com/joshhsoj1902/docker-gogpagent/vendor/github.com/divan/gorilla-xmlrpc/xml/server.go
RUN env
RUN ls -ltr
RUN make build
ENV USERNAME=agent
HEALTHCHECK CMD curl --fail http://localhost:12679/health || exit 1  

CMD ["/go/src/github.com/joshhsoj1902/docker-gogpagent/main"]
