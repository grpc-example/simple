## 准备工作
```bash
brew install protobuf
protoc --version #最新版本
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH=$PATH:$GOPATH/bin
```
## 下载demo
```bash
git clone https://github.com/grpc-example/simple.git
cd simple
go mod tidy
```
## 编译.proto文件命令
```bash
make api
```

## run server
```bash
go run server.go
```

## run client
```bash
go run client.go
```