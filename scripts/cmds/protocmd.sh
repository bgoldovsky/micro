go get -u github.com/golang/protobuf 
go get -u github.com/golang/protobuf/proto 
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc

protoc --go_out=. *.proto
protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/squarer.proto

./evans api/proto/squarer.proto -p 8081
call MyMethod