#安装protoc3
protoc --proto_path=./../third_party --proto_path=./ --bm_out=:./ ./api.proto

protoc --proto_path=./../third_party --proto_path=./ --gofast_out=plugins=grpc:./ ./api.proto

