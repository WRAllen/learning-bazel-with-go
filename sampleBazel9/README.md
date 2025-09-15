# 介绍

基于sampleBazel8的基础，介绍了proto里面如果有引用到第三方的依赖

要如何进行处理和编写

# 运行方式

基于proto自动生成pb.go和_grpc.pb.go的文件
```
make proto
```

把proto和grpc文件移动到api/hello/v1下
```
make move-proto
```



# 项目结构

```
sampleBazel9/
├── api/hello/v1/
│   ├── hello.proto          # 包含新 import 的 proto 文件
│   └── BUILD.bazel          # 配置 proto 依赖
├── BUILD.bazel              # 根构建配置
├── MODULE.bazel             # 管理外部依赖
├── Makefile                 # 构建脚本
└── README.md                # 项目文档
```


## 项目说明

这是一个专注于 **Bazel Proto 和 gRPC 代码生成** 的演示项目，展示如何通过 Bazel 自动处理 proto 文件中的新依赖并生成 gRPC 代码：

下面是用到的第三方依赖
- `google/protobuf/empty.proto`
- `google/protobuf/timestamp.proto`

最重要的就是MODULE.bazel里面添加这行
```
bazel_dep(name = "protobuf", version = "29.1")
```
这个会自动的去下载拉取第三方依赖，如果没有bazel的话，用户就需要去手动下载相关依赖，并且导入到项目里面

后续结合go的web项目，就痛sampleBazel8一样，编写go的部分，然后把grpc里面写好的接口copy到自己的controller层下，进行开发
