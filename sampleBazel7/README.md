# 介绍

如何基于proto文件生成pb文件和对应的grpc文件

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

其实这样和proto没啥关系了，其实就是在bazel里面引用了go的文件

对于api/hello/v1下生成的BUILD文件中的go_proto_library这里compilers使用错误的问题，暂时没找到原因，

我的gazelle版本也挺新的0.43，当前最高才0.44（2025年6月27日）

由于是中间状态的项目，这里也不做深入研究了，重点是后续用bazel自动生成pb文件。

### 如何编写接口

刚开始接触proto的时候，不知道接口怎么写，后面发现可以从xxxx_grpc.pb.go文件里面摘抄。

如下段落

```
type UnimplementedHelloServiceServer struct{}

func (UnimplementedHelloServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
```

在接口层，自定义一个struct继承这个UnimplementedHelloServiceServer就行

然后在http那边注册进去，相关代码

```
// grpc服务创建
grpcServer := grpc.NewServer()
// 注册hello相关接口
api.RegisterHelloServiceServer(grpcServer, &HelloServer{})
```

### 测试接口

推荐一个工具，grpcurl, 可以直接用brew安装

查看当前grpc的服务

```
grpcurl -plaintext localhost:8080 list

api.hello.v1.HelloService
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection\
```

这里注意服务启动是对应的reflection要设置一下

```
reflection.Register(grpcServer)
```

查看服务对应的接口

```
grpcurl -plaintext localhost:8080 list api.hello.v1.HelloService

api.hello.v1.HelloService.Hello
```

测试接口

```
grpcurl -plaintext \
  -d '{"name": "Alice"}' \
  localhost:8080 api.hello.v1.HelloService.Hello

# 返回数据如下
{}

# 服务也输出了
====
```

## 运行结果

运行输出如下

注意这里不用make run，因为需要手动调整一下BUILD文件

```
-> % make only-run
bazel run //:sampleBazel6 -- serve
INFO: Analyzed target //:sampleBazel6 (233 packages loaded, 4665 targets configured).
INFO: From Linking external/abseil-cpp+/absl/debugging/libdebugging_internal.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/debugging/_objs/debugging_internal/elf_mem_image.o has no symbols
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/debugging/_objs/debugging_internal/vdso_support.o has no symbols
INFO: From Linking external/abseil-cpp+/absl/base/libbase.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/base/_objs/base/unscaledcycleclock.o has no symbols
INFO: From Linking external/abseil-cpp+/absl/crc/libcrc32c.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/crc/_objs/crc32c/crc_memcpy_x86_arm_combined.o has no symbols
INFO: From Linking external/abseil-cpp+/absl/synchronization/libsynchronization.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/synchronization/_objs/synchronization/futex_waiter.o has no symbols
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/synchronization/_objs/synchronization/sem_waiter.o has no symbols
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/synchronization/_objs/synchronization/win32_waiter.o has no symbols
INFO: From Linking external/protobuf+/src/google/protobuf/io/libio_win32.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/protobuf+/src/google/protobuf/io/_objs/io_win32/io_win32.o has no symbols
warning: /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: archive library: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/protobuf+/src/google/protobuf/io/libio_win32.a the table of contents is empty (no object file members in the library define global symbols)
INFO: From Linking external/abseil-cpp+/absl/strings/libcord.a [for tool]:
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/libtool: file: bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/abseil-cpp+/absl/strings/_objs/cord/cord_buffer.o has no symbols
INFO: From Linking external/protobuf+/protoc [for tool]:
ld: warning: ignoring duplicate libraries: '-lm', '-lpthread'
INFO: Found 1 target...
Target //:sampleBazel6 up-to-date:
  bazel-bin/sampleBazel6_/sampleBazel6
INFO: Elapsed time: 209.739s, Critical Path: 51.52s
INFO: 925 processes: 161 action cache hit, 420 internal, 505 darwin-sandbox.
INFO: Build completed successfully, 925 total actions
INFO: Running command line: bazel-bin/sampleBazel6_/sampleBazel6 <args omitted>
{"level":"info","time":"2025-06-27 15:22:48.242","caller":"pkg/service.go:44","msg":"Starting gRPC server","port":"8080"}
====
```
