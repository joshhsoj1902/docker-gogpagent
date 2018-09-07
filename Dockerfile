FROM golang:latest@sha256:e8e4c4406217b415c506815d38e3f8ac6e05d0121b19f686c5af7eaadf96f081
RUN mkdir -p /go/src/github.com/joshhsoj1902/docker-gogpagent
ADD . /go/src/github.com/joshhsoj1902/docker-gogpagent/
WORKDIR /go/src/github.com/joshhsoj1902/docker-gogpagent
RUN env
RUN ls -ltr
RUN make build
ENV USERNAME=agent
CMD ["/go/src/github.com/joshhsoj1902/docker-gogpagent/main"]
