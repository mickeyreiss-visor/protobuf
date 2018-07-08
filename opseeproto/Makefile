generate:
	protoc --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. --proto_path=/go/src:. --proto_path=/go/src/github.com/opsee/protobuf/vendor/ *.proto
	protoc --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. --proto_path=/go/src:. --proto_path=/go/src/github.com/opsee/protobuf/vendor/ types/*.proto

.PHONY:
	generate
