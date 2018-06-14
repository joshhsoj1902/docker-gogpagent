FROM golang:latest@sha256:db260e19d31a9c6794d35aae1bf2cd30f1b4db88c3094a18299c10ed02eb4dee
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN make build
CMD ["/app/main"]
