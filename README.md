## 准备工作
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
export PATH=$PATH:$GOPATH/bin
```
## 编译.proto文件命令
```
protoc --go_out=plugins=grpc:. hello.proto
```
## 下载demo
```
git clone https://github.com/grpc-example/simple.git
cd simple
```

## run server
```
go run server.go
```

## run client
```
go run cli.go
```