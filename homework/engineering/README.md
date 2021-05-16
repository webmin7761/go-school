# go 进阶训练营 第四周 作业

## 问题

 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

## 思路

1. 设定业务场景为 - 航空机票价格服务（FareService），提供下面5个接口

    - CreateFare 机票价格规则新增
        - 起始机场、目地机场、开始旅行日期、结束旅行日期、乘客类别（成人、儿童和婴儿）、票价

    - UpdateFare 机票价格规则更新

    - DeleteFare 机票价格规则删除

    - GetFare 机票价格规则按ID获取

    - Pricing 机票价格计算

        - 根据起始机场、目地机场、乘机日期、乘客类别（成人、儿童和婴儿）、条件计算适用票价。

2. 使用Protobuf定义接口协议

3. 采用Kratos生成HTTP和grpc使用框架

4. 使用ent生成实体框架，处理DB操作

5. 使用wrie生成注入代码

## TODO

1. [X] 工程化目录

2. [X] 服务API设计

3. [X] 表结构设计

4. [x] kratos使用

5. [x] ent使用

6. [x] wrie使用

7. [x] 测试

## 参考

- [Go-gRPC 实践指南](https://www.bookstack.cn/read/go-grpc/summary.md)
- [Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases)
- [kratos blog示例](https://github.com/go-kratos/kratos/blob/main/examples/blog)
- [ent文档](https://entgo.io/zh/docs/getting-started)

## 编译

### kratos生成接口

```sh
cd api\fare\v1
kratos proto client fare.proto --proto_path=../../../third_party -I=.
```

### ent

1. 生成结构 Fare 于 `<project>/ent/schema/` 目录内 `ent init Fare`

2. 生成对DB的各种操作`ent generate ./ent/schema`

### 项目编译

```sh
cd cmd\fare
#生成注入代码
wrie
go build github.com/webmin7761/go-school/homework/engineering/cmd/fare
```

## 测试

- [Postman测试用例](test/data/_postman_engineering.json)

## protobuf备忘

```sh
go get -u github.com/golang/protobuf

go get -u github.com/golang/protobuf/protoc-gen-go

go get -u github.com/lazada/protoc-gen-go-http

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-http_out=. --proto_path=d:/protobuf -I=. price.proto
```
