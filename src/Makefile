### Makefile for tcp-proxy

GOPATH ?= $(shell go env GOPATH) 

# Ensure Gopath is set before running build process
ifeq "$(GOPATH)" ""
	$(error Please set the enviroment variable GOPATH before running `make`)
endif 



PREFIX=$(shell pwd)/../  
BINDIR=${PREFEX}/bin
PROJECT  := proxyd 
CUREDIR  := $(shell pwd)
GO       := go 
CPU_COUNT := 2
GOBUILD  := GOPATH=$(CUREDIR):$(GOPATH) CGO_ENABLED=0 $(GO) build -v -x -p $(CPU_COUNT) -o $(PROJECT)   

## other 
ARCH     := "`uname -s`"
LINUX    := "Linux" 
MAC      := "Darwin"


TARGET = ""

build:
	$(GOBUILD) 

## exclude clean file exist
.PHONY: cleanall cleanobj cleandiff clean
.PHONY: install clean all 

clean:
	$(GO) clean -i ./...
	rm -rf $(GOBUILD) 


test:
	@echo 'now --> there is not exist ' 

install: 
	@echo 'now --> start to install ' 

