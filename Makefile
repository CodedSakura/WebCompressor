export PATH := ${HOME}/go/bin/:${HOME}/go/go1.21.0/bin:${PATH}

all: WebCompressor

deps:
	go get

WebCompressor: deps
	go build .

dev:
	CompileDaemon -color -command="./WebCompressor"
