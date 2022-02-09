protoc  --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:"pkg/api/v1" user.proto
protoc  --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:"pkg/api/v1" user-service.proto
protoc  --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:"pkg/api/v1" user-service.proto
