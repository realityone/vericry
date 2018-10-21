gosrc = $(GOPATH)/src
gogosrc = $(gosrc)/github.com/gogo/protobuf
current_dir = $(shell pwd)
proto_dir = $(current_dir)/pkg/proto
schema_dir = $(proto_dir)/schema

proto:
	protoc --gofast_out=$(schema_dir) --proto_path=/usr/local/include:$(gosrc):$(gogosrc):$(schema_dir) $(schema_dir)/*.proto

migrationenv:
	cargo install diesel_cli
