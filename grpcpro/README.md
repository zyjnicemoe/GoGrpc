# 在pbfile下执行
```bash
protoc -I. -I%GOPATH%/include -I%GOPATH%/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.
16.0/third_party/googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. *.proto
```