> 添加依赖
```bash
 go get -t github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
 go get -t github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
 go get -t github.com/golang/protobuf/protoc-gen-go
```

### 在pbfile下执行
```bash
protoc -I. -I%GOPATH%/include -I%GOPATH%/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. *.proto
```