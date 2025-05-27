.PHONY: help setup tearodwn build clean

export GEMINI_API_KEY

help:
	@cat $(firstword $(MAKEFILE_LIST))

setup: \
	dependency \
	dependency/tamakiii/go-genai

teardown:
	rm -rf dependency

build: \
	bin \
	bin/chat

clean:
	rm -rf bin

dependency:
	-mkdir $@

dependency/tamakiii/go-genai:
	git clone git@github.com:tamakiii/go-genai.git $@

bin:
	-mkdir $@

bin/chat: chat.go
	go build -o $@ $<
