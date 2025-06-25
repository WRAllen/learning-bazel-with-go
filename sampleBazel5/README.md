# 介绍

基于sampleBazel4，添加了配置文件相关的内容

介绍了在gazelle自动生成BUILD后，如何把配置文件一起打包

# 运行方式

```
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
    └── service.go

4 directories, 16 files
```

该项目基于sampleBazel4，介绍了如果有和代码不是强相关的文件，比如配置文件，这些如何用bazel打包起来。

同3一样，开始编写根目录MODULE.bazel和BUILD.bazel（这里有疑惑可以看sampleBazel3，里面有详细介绍）

如果我们直接 `make run`会发现可以生成build，但是会报错

```
panic: 加载配置失败: 读取配置失败: Config File "config" Not Found in "[/private/var/tmp/_bazel_xxxx/5f8cd301ddb0885d21d262827a3f974f/execroot/_main/bazel-out/darwin_x86_64-fastbuild/bin/sampleBazel5_/sampleBazel5.runfiles/_main/config]"
```

但是gazelle正常生成了cmd，config，pkg下的BUILD，并且改写了根目录的BUILD（所有的BULID里面除了有标注为手写的，其他的都是gazelle自动生成）

看到错误可以发现时加载配置的时候触发了写的panic，说明Bazel在跑他的时候config相关的yaml文件没有被打包进去

`Config File "config" Not Found in` 这里提示的很明显了，所以我们要加入配置

### 定义使用到的文件资源

当前项目只有config下，有需要用的文件，所以打开config下的BUILD添加如下内容

```
filegroup(
    name = "all_configs",
    srcs = glob(["*.yaml"]),
    visibility = ["//visibility:public"],
)
```

具体这个是干嘛的，都是什么意思，具体注释详见文件

### 在根目录BUILD引用这个资源

在BUILD文件的go_library添加上如下

```
data = ["//config:all_configs"],

```

这里指的是，用config文件下的all_configs资源，具体注释详见文件

这样就ok了，然后我们就能把config下的配置都进行打包了


## 运行结果

运行输出如下

```
-> % make run  
bazel clean --expunge
Starting local Bazel server (8.3.0) and connecting to it...
INFO: Starting clean (this may take a while). Use --async if the clean takes more than several minutes.
bazel run //:gazelle
Starting local Bazel server (8.3.0) and connecting to it...
INFO: Analyzed target //:gazelle (126 packages loaded, 8577 targets configured).
INFO: Found 1 target...
Target //:gazelle up-to-date:
  bazel-bin/gazelle-runner.bash
  bazel-bin/gazelle
INFO: Elapsed time: 92.971s, Critical Path: 56.50s
INFO: 67 processes: 18 internal, 49 darwin-sandbox.
INFO: Build completed successfully, 67 total actions
INFO: Running command line: bazel-bin/gazelle
bazel mod tidy
bazel build //:sampleBazel5
INFO: Analyzed target //:sampleBazel5 (83 packages loaded, 899 targets configured).
INFO: Found 1 target...
Target //:sampleBazel5 up-to-date:
  bazel-bin/sampleBazel5_/sampleBazel5
INFO: Elapsed time: 9.149s, Critical Path: 3.28s
INFO: 56 processes: 4 internal, 52 darwin-sandbox.
INFO: Build completed successfully, 56 total actions
bazel run //:sampleBazel5 -- serve
INFO: Analyzed target //:sampleBazel5 (0 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //:sampleBazel5 up-to-date:
  bazel-bin/sampleBazel5_/sampleBazel5
INFO: Elapsed time: 0.252s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Running command line: bazel-bin/sampleBazel5_/sampleBazel5 <args omitted>
{"level":"info","time":"2025-06-25 14:44:11.561","caller":"pkg/service.go:32","msg":"Starting server","port":"8080"}
```
