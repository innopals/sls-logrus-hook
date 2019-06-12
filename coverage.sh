#!/bin/bash
t=coverage.tmp
go test -coverprofile=$t -bench=. $@ && go tool cover -func=$t | grep -v log.pb.go && unlink $t
