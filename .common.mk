## protoc
PROTOC ?= protoc

PROTOC_INCLUDE ?= -I . \
-I vendor/github.com/envoyproxy/protoc-gen-validate \

PROTOC_PLUGINS ?= --go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
--grpc-gateway_out=. \
--grpc-gateway_opt=logtostderr=true,paths=source_relative \
--validate_out=. \
--validate_opt=paths=source_relative,lang=go

PROTOC_TAGGER ?= --gotag_out=. \
--gotag_opt=xxx="bson+\"-\"",output_path
