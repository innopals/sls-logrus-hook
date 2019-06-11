#!/bin/bash
t=coverage.tmp
go test -coverprofile=$t $@ && go tool cover -func=$t && unlink $t
