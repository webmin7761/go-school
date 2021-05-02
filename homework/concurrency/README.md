# go 进阶训练营 第三周 作业

## 问题

 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

## 思路

1. 优雅关闭

2. 何时以及怎么关闭http server

## 参考

`https://github.com/da440dil/go-workgroup/blob/master/template/signal/signal.go`

`https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-sync-primitives/#errgroup`

## 编译

`go run github.com/webmin7761/go-school/homework/concurrency`

## 使用

- <http://localhost:8080/>
- <http://127.0.0.1:6060/debug/pprof/>
