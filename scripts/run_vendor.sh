#!/bin/bash

rm -rf vendor/
mkdir -p vendor
go mod tidy
go mod vendor
