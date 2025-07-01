# 介绍

基于sampleBazel7的自动生成pb/gprc，结合6的web项目

说明如何做到先写proto再写接口服务，然后build项目。

# 运行方式

详见Makefile文件

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
│           ├── hello.pb.go	 # 后续生成
│           ├── hello.proto
│           └── hello_grpc.pb.go # 后续生成
├── cmd
│   ├── BUILD.bazel
│   ├── root.go
│   └── serve.go
├── config
│   ├── BUILD.bazel
│   ├── config.go
│   └── config.yaml
├── copy_proto_outputs.sh
├── go.mod
├── go.sum
├── main.go
└── pkg
    ├── BUILD.bazel
    ├── controller.go
    └── service.go
```

流程如下

### 定义整体架构

首先编写sample6的整体框架，然后先不写和接口相关的部分，比如先创建个空的controller，service里面启动服务时先不注册接口

保证这个项目可以通过 `go run main.go serve`进行启动

然后编写proto文件，基于proto生成grpc代码，在服务里面继承实现这个接口，启动服务时注册对应的接口即可

# 运行过程

### 生成BUILD

运行gazelle生成BUILD文件，手动调整生成的BUILD文件

运行 `make generate 去生成BUILD`

### 生成pb/grpc代码

然后通过bazel去自动生成pb/grpc文件，跑脚本把生成的文件加入api/hello/v1下

先 make build-proto运行pb/grpc生成

```
-> % make build-proto
bazel build //api/hello/v1:hello
INFO: Analyzed target //api/hello/v1:hello (186 packages loaded, 4445 targets configured).
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
Target //api/hello/v1:hello up-to-date:
  bazel-bin/api/hello/v1/hello.x
INFO: Elapsed time: 213.041s, Critical Path: 67.72s
INFO: 1033 processes: 2 action cache hit, 420 internal, 613 darwin-sandbox.
INFO: Build completed successfully, 1033 total actions
```

运行 `make move-proto`移动到api下

```
-> % make move-proto   
./copy_proto_outputs.sh
🔍 Searching for .pb.go files under bazel-bin/api/hello/v1...
📄 Copying files to api/hello/v1:
 → hello.pb.go
bazel-bin/api/hello/v1/api_go_proto_pb_/github.com/wrallen/sampleBazel8/api/hello/v1/hello.pb.go -> api/hello/v1/hello.pb.go
 → hello_grpc.pb.go
bazel-bin/api/hello/v1/api_go_proto_/github.com/wrallen/sampleBazel8/api/hello/v1/hello_grpc.pb.go -> api/hello/v1/hello_grpc.pb.go
✅ Done.
```

## 继承编写接口

通过从hello_grpc.pb.go里面写好的空接口，然后集成到自己的服务里面，我这里就是把定义的接口，在controller.go里面继承实现了一下

## 测试go是否能直接运行

执行命令 `make go-run`或者如下命令

```
-> % go run main.go serve
{"level":"info","time":"2025-07-01 15:47:33.399","caller":"pkg/service.go:44","msg":"Starting gRPC server","port":"8080"}
```

发现是ok的

## 在pkg里面引用api

这个时候如果直接用bazel build的话会出现如下问题

```
-> % make build  
bazel build //:sampleBazel8
INFO: Analyzed target //:sampleBazel8 (51 packages loaded, 324 targets configured).
ERROR: /Users/wangyu2/Work/learning-bazel-with-go/sampleBazel8/pkg/BUILD.bazel:3:11: GoCompilePkg pkg/pkg.a failed: (Exit 1): builder failed: error executing GoCompilePkg command (from target //pkg:pkg) bazel-out/darwin_x86_64-opt-exec-ST-d57f47055a04/bin/external/rules_go++go_sdk+main___download_0/builder_reset/builder compilepkg -sdk external/rules_go++go_sdk+main___download_0 -goroot ... (remaining 35 arguments skipped)

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
compilepkg: missing strict dependencies:
        /private/var/tmp/_bazel_wangyu2/fa29c67326d5290fa02d0e6c6d37a8f1/sandbox/darwin-sandbox/712/execroot/_main/pkg/controller.go: import of "github.com/wrallen/sampleBazel8/api/hello/v1"
        /private/var/tmp/_bazel_wangyu2/fa29c67326d5290fa02d0e6c6d37a8f1/sandbox/darwin-sandbox/712/execroot/_main/pkg/service.go: import of "github.com/wrallen/sampleBazel8/api/hello/v1"
No dependencies were provided.
Check that imports in Go sources match importpath attributes in deps.
Target //:sampleBazel8 failed to build
Use --verbose_failures to see the command lines of failed build steps.
INFO: Elapsed time: 6.437s, Critical Path: 2.99s
INFO: 54 processes: 5 internal, 49 darwin-sandbox.
ERROR: Build did NOT complete successfully
make: *** [build] Error 1
```

可以发现是引用出了问题，其实就是pkg那边用api这边的定义的时候bazel不知道是啥东西，因为pkg里面用的是后来生成的文件，

这里我们再次运行 `make generate` 的话会覆盖掉之前api生成proto相关的部分，这里不用gazelle去自动更新BUILD，这里我们手动在pkg的BUILD的deps添加如下内容即可

```
"//api/hello/v1:hello"
```

## 使用bazel进行打包

```
-> % make build   
bazel build //:sampleBazel8
INFO: Analyzed target //:sampleBazel8 (1 packages loaded, 484 targets configured).
INFO: Found 1 target...
Target //:sampleBazel8 up-to-date:
  bazel-bin/sampleBazel8_/sampleBazel8
INFO: Elapsed time: 0.886s, Critical Path: 0.31s
INFO: 3 processes: 1086 action cache hit, 2 internal, 1 darwin-sandbox.
INFO: Build completed successfully, 3 total actions
```

## 运行服务

```
-> % make run  
bazel run //:sampleBazel8 -- serve
INFO: Analyzed target //:sampleBazel8 (0 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //:sampleBazel8 up-to-date:
  bazel-bin/sampleBazel8_/sampleBazel8
INFO: Elapsed time: 0.238s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Running command line: bazel-bin/sampleBazel8_/sampleBazel8 <args omitted>
{"level":"info","time":"2025-07-01 16:05:10.401","caller":"pkg/service.go:44","msg":"Starting gRPC server","port":"8080"}
```
