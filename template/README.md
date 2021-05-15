# README

## protoc
```sh
cd proto
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=template/proto/google/api:. ./myapp/user.proto

protoc --grpc-gateway_out=logtostderr=true:. ./myapp/user.proto
```

## 证书

https://blog.csdn.net/weixin_40280629/article/details/113563351

## 打包命令
```sh
go build -o ./cmd/myapp/main ./cmd/myapp/
```