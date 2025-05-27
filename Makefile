.PHONY: help build

export GEMINI_API_KEY

help:
	@cat $(firstword $(MAKEFILE_LIST))

build:
	go build
