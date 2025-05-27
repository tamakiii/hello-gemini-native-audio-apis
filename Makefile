.PHONY: help build

export GEMINI_API_KEY

help:
	@cat $(firstword $(MAKEFILE_LIST))

build: \
	bin \
	bin/chat

clean:
	rm -rf bin

bin:
	-mkdir $@

bin/chat: chat.go
	go build -o $@ $<
