cd pbfiles && protoc --go_out=plugins=grpc,paths=source_relative:../services  Prod.proto
 protoc --go_out=plugins=grpc,paths=source_relative:../services   Orders.proto
protoc --go_out=plugins=grpc,paths=source_relative:../services   Users.proto

protoc --go_out=plugins=grpc,paths=source_relative:../services --validate_out=lang=go,paths=source_relative:../services   Models.proto

protoc  --grpc-gateway_out=logtostderr=true,paths=source_relative:../services Prod.proto
protoc  --grpc-gateway_out=logtostderr=true,paths=source_relative:../services Orders.proto
protoc  --grpc-gateway_out=logtostderr=true,paths=source_relative:../services Users.proto
cd ..