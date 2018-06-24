build-docker:
	docker build -t joshhsoj1902/docker-gogpagent .

build:
	go build -o main .

start:
	docker run joshhsoj1902/docker-gogpagent

.PHONY: build
