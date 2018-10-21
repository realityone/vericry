gosrc = $(GOPATH)/src
gogosrc = $(gosrc)/github.com/gogo/protobuf
current_dir = $(shell pwd)
proto_dir = $(current_dir)/pkg/proto
model_dir = $(proto_dir)/model

proto:
	protoc --gofast_out=$(model_dir) --proto_path=/usr/local/include:$(gosrc):$(gogosrc):$(model_dir) $(model_dir)/*.proto

migrationenv:
	cargo install diesel_cli
