.PHONY: default bindata

BINDATA := $(shell command -v go-bindata 2> /dev/null)

bindata:
ifndef BINDATA
	go get -u github.com/jteeuwen/go-bindata/...
endif
	@go-bindata -pkg utils -o lib/utils/resources.go resources/...