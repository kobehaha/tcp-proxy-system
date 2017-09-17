#!/bin/bash

CURDIR=$(pwd)

GO=$(which go) 
echo "fmt all go file"

GO fmt ./*.go







