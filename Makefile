gosrc = $(GOPATH)/src
gogosrc = $(gosrc)/github.com/gogo/protobuf
current_dir = $(shell pwd)
proto_dir = $(current_dir)/pkg/proto

proto:
	protoc --gofast_out=$(proto_dir) --proto_path=/usr/local/include:$(gosrc):$(gogosrc):$(proto_dir) $(proto_dir)/*.proto
