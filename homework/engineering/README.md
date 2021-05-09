# go 进阶训练营 第四周 作业

## 问题

 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

## 思路

 设定业务场景为 - 航空机票价格计算服务

 根据起始机场、目地机场、乘机日期、乘客类别（成人、儿童和婴儿）、舱等(公务舱、经济舱)条件计算适用票价以及相对应的退票、变更条件。

## TODO

1. [] 工程化目录

2. [] 服务API设计

3. [] 表结构设计

4. [] 业务层设计及实现

5. [] 服务脚手架实现

## 参考

## 编译

`go run github.com/webmin7761/go-school/homework/engineering`

## 使用

## grpc备忘

- [Go-gRPC 实践指南](https://www.bookstack.cn/read/go-grpc/summary.md)

- [Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases)

```sh
go get -u github.com/golang/protobuf

go get -u github.com/golang/protobuf/protoc-gen-go

go get -u github.com/golang/protobuf/protoc-gen-go-grpc

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=d:/protobuf -I=. price.proto
```
