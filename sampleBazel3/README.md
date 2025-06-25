# 介绍

这是一个多层次的有依赖第三方包的go的web项目

介绍了基于gazelle如何去自动更新BUILD.bazel文件

# 运行方式

```
# 清除Bazel生成的文件
make clean

# 使用gazelle生成BUILD文件（这里其实已经生成好了）
make build

# 使用Bazel进行打包
make build

# 使用Bazel运行程序，这里run已经做了上面的clean和build了
make run
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
├── go.mod
├── go.sum
├── main.go
└── pkg
    ├── BUILD.bazel
    └── controller.go

2 directories, 10 files
```

该项目主要介绍了如果有第三方的包的依赖，如何在Bazel里面做好，sampleBazel2由于用的都是go内置的包，所以不涉及。该web项目使用了第三方的zap包，替代2里面和日志相关的输出

这里zap只是作为一个例子，其他的依赖也是同理。

### 定义MODULE.bazel

最重要是在MODULE里面定义好gazelle相关的部分，下面摘出gazelle相关代码，具体说明注释请到对应文件里面阅读

```
bazel_dep(name = "gazelle", version = "0.43.0")

go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")

use_repo(go_deps, "org_uber_go_zap")

```

由于我也在学习阶段，这里被GPT坑了几次，相关的坑也写在MODULE.bazel文件里面了，这里还知道了一个bazel的命令

```
bazel mod tidy
```

这个同go的tidy，他会去整理文件会 **解析并检查你的 `MODULE.bazel` 文件，下载依赖，并更新 `.lock` 文件** ，以确保 Bazel 构建环境中的所有依赖是最小且一致的。

### 定义BUILD.bazel

MODULE.bazel文件写完后，其他的BUILD.bazel文件就简单了，如下

```
load("@gazelle//:def.bzl", "gazelle")

gazelle(name = "gazelle")

```

只需要在你每个目录下面创建一个BUILD文件，写上上面2行

然后再运行如下命令

```
# 或者是Makefile里面写好的make generate
bazel run //:gazelle

```

他就会自动的去更新对应BUILD文件里面的内容，帮你编写BUILD文件需要的一切，（sampleBazel4里介绍如何自动生成BUILD）

比如go_library这些，他都会自动去生成，会把需要用到的第三方依赖自动写到对应BUILD里面的deps里，

具体也是见根目录的BUILD和pkg下面的BUILD。

### 更新MODULE里的use_repo依赖

上面BUILD自动生成后，然后就摘抄deps里面的依赖更新到MODULE里面的use_repo下，

这里注意，比如某个BUILD里面的deps如下

```
@org_uber_go_zap//:zap
```

写到use_repo里面时只需要

```
org_uber_go_zap
```

### 另一种更新MODULE的依赖方式

当然也有一种办法能不用一个一个去子目录的BUILD里面摘抄deps，直接写

```
# go_deps后面啥也不写
use_repo(go_deps)
```

然后运行 `bazel run //:gazelle `就会出现如下的提示

```
Not imported, but reported as direct dependencies by the extension (may cause the build to fail):
   org_uber_go_zap
```

把这些摘抄到use_repo里面就行（sampleBazel4里面介绍如何自动生成相关依赖）


至此就完成了有第三方依赖的项目

## 运行结果

运行输出如下

```
-> % make run  
bazel clean --expunge
INFO: Starting clean (this may take a while). Use --async if the clean takes more than several minutes.
bazel build //:sampleBazel3
Starting local Bazel server (8.3.0) and connecting to it...
INFO: Analyzed target //:sampleBazel3 (97 packages loaded, 6058 targets configured).
INFO: Found 1 target...
Target //:sampleBazel3 up-to-date:
  bazel-bin/sampleBazel3_/sampleBazel3
INFO: Elapsed time: 69.904s, Critical Path: 36.33s
INFO: 21 processes: 6 internal, 15 darwin-sandbox.
INFO: Build completed successfully, 21 total actions
bazel run //:sampleBazel3
INFO: Analyzed target //:sampleBazel3 (32 packages loaded, 316 targets configured).
INFO: Found 1 target...
Target //:sampleBazel3 up-to-date:
  bazel-bin/sampleBazel3_/sampleBazel3
INFO: Elapsed time: 0.627s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Running command line: bazel-bin/sampleBazel3_/sampleBazel3
{"level":"info","ts":1750763064.1256928,"caller":"pkg/controller.go:25","msg":"Starting server","port":"8080"}
```
