generate:
	protoc --gogoopsee_out=plugins=grpc+graphql+bitflags,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:flags --proto_path=/go/src:flags --proto_path=/go/src/github.com/opsee/protobuf/vendor/ flags/*.proto
	protoc --gogoopsee_out=plugins=grpc+graphql,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:dessert --proto_path=/go/src:dessert --proto_path=/go/src/github.com/opsee/protobuf/vendor/ dessert/*.proto
	protoc --gogoopsee_out=plugins=grpc+graphql,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. --proto_path=/go/src:. --proto_path=/go/src/github.com/opsee/protobuf/vendor/ *.proto

clean:
	rm -f {,**/}*.pb.go
	rm -f {,**/}*pb_test.go

.PHONY: clean generate
