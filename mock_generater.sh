#!/bin/bash

export mockery=$GOPATH/bin/mockery

mockery --recursive=true --all --output=./_mocks/ --outpkg=service
