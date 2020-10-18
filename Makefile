-include .common.mk

export

deps:
	go mod vendor

generate:
	go generate ./...

proto:
	$(PROTOC) \
		$(PROTOC_INCLUDE) \
		$(PROTOC_PLUGINS) \
		`ls *.proto`

test:
	go test `go list ./... | grep -v /tools`

tool-chain:
	go get -u \
		github.com/envoyproxy/protoc-gen-validate

.DEFAULT: deps generate proto
