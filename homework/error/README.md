# go 进阶训练营 第二周 作业

## 问题

 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层

## 思路

dao中的操作是按具体对象来操作的，sql.ErrNoRows是通用说法，并不是表意的，
例如：
如果是按用户名查找用户的场景下，查找不到用户应该把sql.ErrNoRows Wrap到“用户不存的”错误里
这样在上层如果要进行某种通用型处理也可以兼顾的到

## 代码文件说明

### init.go

1. 准备测试表和数据(Sqlite)

2. 初始化Conn

### dao.go

1. 操作User的Dao

2. FindByName,QueryBySex

3. sql.ErrNoRows Wrap用法

### main.go

> 使用例子
> `go run github.com/webmin7761/go-school/homework/error`
