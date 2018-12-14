FROM golang:latest@sha256:8e4c5d521e884363bf9581f2f4460525d010e6a9aca7e5d4e6ddcb4ade01de44
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
