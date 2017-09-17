### Makefile for tcp-proxy

GOPATH ?= $(shell go env GOPATH) 

# Ensure Gopath is set before running build process
ifeq "$(GOPATH)" ""
	$(error Please set the enviroment variable GOPATH before running `make`)
endif 

CUREDIR  := $(shell pwd)
GO       := GO
GOBUILD  := GOPATH=$(CUREDIR):$(GOPATH) CGO_ENABLED=0 $(GO) build

## other 
ARCH     := "`uname -s`"
LINUX    := "Linux" 
MAC      := "Darwin"


TARGET = ""

build:
	$(GOBUILD) 

clean:
	$(GO) clean -i ./...
	rm -rf *.out




