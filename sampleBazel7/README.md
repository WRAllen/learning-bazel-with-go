# 介绍

介绍通过bazel基于proto文件生成pb，grpc文件

# 运行方式

```

# 清除Bazel生成的文件
make clean

# 使用gazelle更新/创建BUILD文件
make generate

# 进行文件生成
make build

```

# 具体说明

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
│           └── hello.proto
└── go.mod

4 directories, 8 files
```

该项目基于一个空的go的项目，只编写了基础的hello.proto文件，介绍如何用bazel去自动的生成pb文件和对应的grpc文件

因为一般我们编写一个go的web项目，都是先去写proto文件，然后通过protoc去自动生成pb和grpc文件（可以参考sampleBazel6）

这里简单介绍一下pb文件和grpc文件

pb文件：message 结构体定义 + 编解码

grpc文件：gRPC 客户端 & 服务端接口定义

grpc文件是基于pb文件的，因为他里面用到了pb定义的结构体等

## 运行结果

运行输出如下

```
-> % bazel build //api/hello/v1:hello
Starting local Bazel server (8.3.0) and connecting to it...
INFO: Analyzed target //api/hello/v1:hello (269 packages loaded, 11626 targets configured).
INFO: From Linking external/abseil-cpp+/absl/debugging/libdebugging_internal.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/debugging/_objs/debugging_internal/elf_mem_image.o has no symbols
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/debugging/_objs/debugging_internal/vdso_support.o has no symbols
INFO: From Linking external/abseil-cpp+/absl/base/libbase.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/base/_objs/base/unscaledcycleclock.o has no symbols
INFO: From Linking external/abseil-cpp+/absl/synchronization/libsynchronization.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/synchronization/_objs/synchronization/futex_waiter.o has no symbols
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/synchronization/_objs/synchronization/sem_waiter.o has no symbols
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/synchronization/_objs/synchronization/win32_waiter.o has no symbols
INFO: From Linking external/abseil-cpp+/absl/crc/libcrc32c.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/crc/_objs/crc32c/crc_memcpy_x86_arm_combined.o has no symbols
INFO: From Linking external/abseil-cpp+/absl/strings/libcord.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/strings/_objs/cord/cord_buffer.o has no symbols
INFO: From Linking external/protobuf+/src/google/protobuf/io/libio_win32.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/protobuf+/src/google/protobuf/io/_objs/io_win32/io_win32.o has no symbols
warning: /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: archive library: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/protobuf+/src/google/protobuf/io/libio_win32.a the table of contents is empty (no object file members in the library define global symbols)
INFO: From Linking external/protobuf+/protoc [for tool]:
ld: warning: ignoring duplicate libraries: '-lm', '-lpthread'
INFO: Found 1 target...
Target //api/hello/v1:hello up-to-date:
  bazel-bin/api/hello/v1/hello.x
INFO: Elapsed time: 271.708s, Critical Path: 96.81s
INFO: 1033 processes: 422 internal, 611 darwin-sandbox.
INFO: Build completed successfully, 1033 total actions
```

### 生成文件

在bazel-out下的 `项目目录/bazel-out/darwin_x86_64-fastbuild/bin/api/hello/v1`

这里的darwin由于我是mac，不用的机子这里也是不同的

有如下文件

```
-> % tree
.
├── api_go_proto_
│   └── github.com
│       └── wrallen
│           └── sampleBazel7
│               └── api
│                   └── hello
│                       └── v1
│                           └── hello_grpc.pb.go
├── api_go_proto_pb_
│   └── github.com
│       └── wrallen
│           └── sampleBazel7
│               └── api
│                   └── hello
│                       └── v1
│                           └── hello.pb.go
├── api_proto-descriptor-set.proto.bin
├── hello.a
└── hello.x

15 directories, 5 files
```

可以看到生成了hello.pb.go和hello_grpc.pb.go


## gazelle使用说明

在测试这个项目的时候，通过把MODULE文件里面的bazel_dep给注释了后，在通过gazelle去自动生成api下的BUILD发现

gazelle生成的BUILD文件最上面的load，比如

`load("@rules_proto//proto:defs.bzl","proto_library")`

其实和MODULE里面写没写如下内容没有关系

`bazel_dep(name ="rules_proto", version ="7.1.0")`

但是后续 `bazel build`时如果MODULE里面没有对应依赖就会报错


由于我本地也有protoc，其实可以不使用拓展里面的protoc，那简化后的MODULE文件如下

```
module(
    name = "sample_bazel_7",
    version = "1.0",
)

bazel_dep(name = "rules_go", version = "0.55.1")
bazel_dep(name = "gazelle", version = "0.44.0")
# 添加和proto相关的
bazel_dep(name = "rules_proto", version = "7.1.0")

go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")

```
