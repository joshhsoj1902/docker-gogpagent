FROM golang:latest@sha256:4e7490c7830da89b0b573135eea3329ef0645c11f95174fccb282a81516483cd
RUN mkdir -p /go/src/github.com/joshhsoj1902/docker-gogpagent
ADD . /go/src/github.com/joshhsoj1902/docker-gogpagent/
WORKDIR /go/src/github.com/joshhsoj1902/docker-gogpagent
RUN env
RUN ls -ltr
RUN make build
CMD ["/go/src/github.com/joshhsoj1902/docker-gogpagent/main"]
