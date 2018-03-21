#!/bin/bash

export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:/usr/local/go-pach
GOBIN=/usr/local/go/bin/go
echo "fmt all go file"

$GOBIN fmt */*.go







