# 介绍

基于sampleBazel9的基础，使用gazelle去自动生成BUILD文件

# 运行方式
 
运行gazelle自动生成api下的build文件
```
make gazelle
```

生成的文件如下
```python
load("@rules_go//go:def.bzl", "go_library")
load("@rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "api_proto",
    srcs = ["hello.proto"],
    visibility = ["//visibility:public"],
    deps = [
        "//google/api:api_proto",
        "@protobuf//:empty_proto",
        "@protobuf//:timestamp_proto",
    ],
)

go_proto_library(
    name = "api_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc_v2"],
    importpath = "github.com/wrallen/sampleBazel10/api/hello/v1",
    proto = ":api_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//google/api:annotations_proto",
        "//google/api:field_behavior_proto",
    ],
)

go_library(
    name = "hello",
    embed = [":api_go_proto"],
    importpath = "github.com/wrallen/sampleBazel10/api/hello/v1",
    visibility = ["//visibility:public"],
)

```
这里还是之前的那个问题，先把go_proto_library里面的io_bazel_rules_go改为rules_go
原因：@io_bazel_rules_go - 这是旧版本的 rules_go 仓库名称

然后再次运行make proto是会报错，错误如下
```shell
bazel build //api/hello/v1:hello
ERROR: no such package 'google/api': BUILD file not found in any of the following directories. Add a BUILD file to a directory to mark it as a package.
 - google/api
ERROR: /Users/wangyu2/Work/learning-bazel-with-go/sampleBazel10/api/hello/v1/BUILD.bazel:16:17: no such package 'google/api': BUILD file not found in any of the following directories. Add a BUILD file to a directory to mark it as a package.
 - google/api and referenced by '//api/hello/v1:api_go_proto'
ERROR: Analysis of target '//api/hello/v1:hello' failed; build aborted: Analysis failed
INFO: Elapsed time: 0.291s, Critical Path: 0.01s
INFO: 1 process: 1 internal.
ERROR: Build did NOT compl
```
这是因为bazel把我们proto依赖的第三方包当作是本地的包了，对比sampleBazel9，是这样写的
```python
# 忽略其他

proto_library(
    name = "api_proto",
    srcs = ["hello.proto"],
    visibility = ["//visibility:public"],
    deps = [
        "@googleapis//google/api:annotations_proto",
        "@googleapis//google/api:field_behavior_proto",
        "@protobuf//:empty_proto",
        "@protobuf//:timestamp_proto",
    ],
)
# 忽略其他
```
可以发现protobuf和googleapis都是被当作外部依赖处理了，


这里为啥protobuf在gazelle下也能正常被当作是外部依赖呢
因为protobug在MODULE.bazel里面声明了
```python
protoc.toolchain(
    google_protobuf = "com_google_protobuf",
    version = "v29.3",
)
```
所以gazelle在编写protobuf的时候知道他是第三方依赖，但是googleapis只是声明了引用
```python
bazel_dep(name = "googleapis", version = "0.0.0-20250826-a92cee39")
```
