.PHONY: help build

export GEMINI_API_KEY

help:
	@cat $(firstword $(MAKEFILE_LIST))

build: \
	bin \
	bin/chat \
	bin/test_models

clean:
	rm -rf bin

bin:
	-mkdir $@

bin/chat: chat.go
	go build -o $@ $<

bin/test_models: test_models.go
	go build -o $@ $<
