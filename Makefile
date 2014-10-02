# tecna
# https://github.com/chrisenytc/tecna
#
# Copyright (c) 2014, Christopher EnyTC
# Licensed under the MIT license.


test:
	go test -v ./test

start:
	go run server.go

.PHONY: test