# 介绍

基于sampleBazel5，添加了proto相关的部分(暂时手动生成pb.go)

介绍了有proto文件时如何用bazel进行打包

# 运行方式

```
# install和protoc相关的依赖
make init

# 生成proto文件
make generate-proto

# 清除Bazel生成的文件
make clean

# 使用gazelle更新/创建BUILD文件
make generate

# 进行打包
make build

# 使用Bazel运行程序
make run

# 使用build好的可执行文件直接跑
make only-run

# 直接用go跑
make go-run
```

# 具体说明

下面是对改项目的具体说明

## 项目结构

结构如下：

```
-> % tree
.
├── BUILD.bazel
├── MODULE.bazel
├── MODULE.bazel.lock
├── Makefile
├── README.md
├── api
│   └── hello
│       └── v1
│           ├── BUILD.bazel
│           ├── hello.pb.go
│           ├── hello.proto
│           └── hello_grpc.pb.go
├── cmd
│   ├── BUILD.bazel
│   ├── root.go
│   └── serve.go
├── config
│   ├── BUILD.bazel
│   ├── config.go
│   └── config.yaml
├── go.mod
├── go.sum
├── main.go
└── pkg
    ├── BUILD.bazel
    ├── controller.go
    └── service.go

7 directories, 21 files
```

该项目基于sampleBazel5，介绍了如何结合proto文件，当前这个项目的pb.go是手动生成的，后续介绍如何把这一步也用bazel来搞

api/hello/v1下的hello.proto为手动编写的proto3的文件，通过protoc（make init安装）后，自动生成了hello.pb.go,以及hello_grpc.pb.go这2个文件



## 运行结果

运行输出如下

```

```
