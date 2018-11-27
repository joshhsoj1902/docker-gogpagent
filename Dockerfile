FROM golang:latest@sha256:41b92066537ec44eb6b79d67e6cfaa2f1085ceba5bb5257f325e4fd3338888a0
RUN mkdir -p /go/src/github.com/joshhsoj1902/docker-gogpagent
ADD . /go/src/github.com/joshhsoj1902/docker-gogpagent/
WORKDIR /go/src/github.com/joshhsoj1902/docker-gogpagent
RUN cp /go/src/github.com/joshhsoj1902/docker-gogpagent/overrides/server.go /go/src/github.com/joshhsoj1902/docker-gogpagent/vendor/github.com/divan/gorilla-xmlrpc/xml/server.go
RUN env
RUN ls -ltr
RUN make build
ENV USERNAME=agent
CMD ["/go/src/github.com/joshhsoj1902/docker-gogpagent/main"]
