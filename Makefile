gosrc = $(GOPATH)/src
gogosrc = $(gosrc)/github.com/gogo/protobuf
current_dir = $(shell pwd)
proto_dir = $(current_dir)/pkg/proto
model_dir = $(proto_dir)/model
api_dir = $(proto_dir)/api/v1

proto:
	protoc --gofast_out=$(model_dir) --proto_path=/usr/local/include:$(gosrc):$(gogosrc):$(model_dir) $(model_dir)/*.proto
	protoc --gofast_out=plugins=grpc:$(api_dir) --proto_path=/usr/local/include:$(gosrc):$(gogosrc):$(model_dir):$(api_dir) $(api_dir)/*.proto

migrationenv:
	cargo install diesel_cli
