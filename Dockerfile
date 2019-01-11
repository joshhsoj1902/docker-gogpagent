FROM golang:latest@sha256:0c7ba0b5eff462068cb8c16d94248ed44463c30418fac91ee46283e34f483547
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
