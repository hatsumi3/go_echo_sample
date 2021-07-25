#!/usr/bin/env bash

go test -covermode=count -coverprofile=cov/c.out
go tool cover -html=cov/c.out -o cov/coverage.html
