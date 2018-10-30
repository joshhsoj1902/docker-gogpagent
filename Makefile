build-docker:
	docker build -f build/Dockerfile -t joshhsoj1902/docker-gogpagent .

build:
	go build -o main cmd/main/main.go

start:
	docker run joshhsoj1902/docker-gogpagent

.PHONY: build
