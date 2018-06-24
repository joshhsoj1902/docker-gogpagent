FROM golang:latest@sha256:db260e19d31a9c6794d35aae1bf2cd30f1b4db88c3094a18299c10ed02eb4dee
RUN mkdir -p /go/src/github.com/joshhsoj1902/docker-gogpagent
ADD . /go/src/github.com/joshhsoj1902/docker-gogpagent/
WORKDIR /go/src/github.com/joshhsoj1902/docker-gogpagent
RUN env
RUN ls -ltr
RUN make build
CMD ["/go/src/github.com/joshhsoj1902/docker-gogpagent/main"]
