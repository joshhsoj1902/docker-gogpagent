build-docker:
	docker build -t joshhsoj1902/docker-gogpagent .

build:
	go get github.com/gorilla/rpc \
  && go get github.com/divan/gorilla-xmlrpc/xml \
	&& go get github.com/gorilla/handlers \
	&& go build -o main .

.PHONY: build
