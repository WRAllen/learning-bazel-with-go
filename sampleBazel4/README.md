# 介绍

基于sampleBazel3，添加了cobra让项目变成一个cli项目，添加了新的文件夹

介绍了通过gazelle如何去自动生成BUILD.bazel文件，以及如何自动更新依赖

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
├── cmd
│   ├── BUILD.bazel
│   ├── root.go
│   └── serve.go
├── go.mod
├── go.sum
├── main.go
└── pkg
    ├── BUILD.bazel
    └── service.go

3 directories, 13 files
```

该项目基于sampleBazel3，在编写3的时候就发现使用bazel遇到两个不合理的地方

1.需要手动去每个子项目里面写同样的两行代码

2.以及根目录BUILD.bazel里面的use_repo的更新感觉十分的麻烦。

其实是可以做到自动化的，我们只需要按照下面步骤，即可实现自动化

### 定义MODULE.bazel文件

内容如下:

```
module(
    name = "sample_bazel_4",
    version = "1.0",
)

bazel_dep(name = "rules_go", version = "0.54.1")
bazel_dep(name = "gazelle", version = "0.43.0")

my_go_sdk = use_extension("@rules_go//go:extensions.bzl", "go_sdk")
my_go_sdk.download(version = "1.24.2")

go_deps = use_extension("@gazelle//:extensions.bzl", "go_deps")
go_deps.from_file(go_mod = "//:go.mod")
use_repo(
    go_deps
)

```

其他的都同sampleBazel3，只是use_repo这里我们不写具体的依赖

### 定义根目录的BUILD.bazel文件

内容如下:

```
load("@gazelle//:def.bzl", "gazelle")

gazelle(name = "gazelle")

```

还是一样写上面两行，然后运行如下命令去自动生成BUILD

```
bazel run //:gazelle
```

可以看到有如下的warn提示

```
WARNING: /Users/wangyu2/Work/learning-bazel-with-go/sampleBazel4/MODULE.bazel:12:24: The module extension go_deps defined in @gazelle//:extensions.bzl reported incorrect imports of repositories via use_repo():

Not imported, but reported as direct dependencies by the extension (may cause the build to fail):
    com_github_spf13_cobra, org_uber_go_zap

Fix the use_repo calls by running 'bazel mod tidy'.
```

这里sampleBazel3之前也介绍过了，其实他也提示了可以用 `bazel mod tidy`去fix，只是当时没发现。

运行后，会发现子目录cmd和pkg里面自动生成了BUILD文件，但是use_repo这里还是空的，这个时候直接去BUILD就会报错

所以我们运行tidy命令，就会发现他自动的去填充了use_repo需要的依赖了。

至此就完成了自动生成BUILD，自动导入基于go.mod的依赖

## 运行结果

运行输出如下

```
-> % make run
bazel clean --expunge
INFO: Starting clean (this may take a while). Use --async if the clean takes more than several minutes.
bazel run //:gazelle
Starting local Bazel server (8.3.0) and connecting to it...
INFO: Analyzed target //:gazelle (126 packages loaded, 8577 targets configured).
INFO: Found 1 target...
Target //:gazelle up-to-date:
  bazel-bin/gazelle-runner.bash
  bazel-bin/gazelle
INFO: Elapsed time: 89.906s, Critical Path: 53.94s
INFO: 67 processes: 18 internal, 49 darwin-sandbox.
INFO: Build completed successfully, 67 total actions
INFO: Running command line: bazel-bin/gazelle
bazel mod tidy
bazel build //:sampleBazel4
INFO: Analyzed target //:sampleBazel4 (46 packages loaded, 440 targets configured).
INFO: Found 1 target...
Target //:sampleBazel4 up-to-date:
  bazel-bin/sampleBazel4_/sampleBazel4
INFO: Elapsed time: 3.945s, Critical Path: 2.12s
INFO: 20 processes: 4 internal, 16 darwin-sandbox.
INFO: Build completed successfully, 20 total actions
bazel run //:sampleBazel4 -- serve
INFO: Analyzed target //:sampleBazel4 (0 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //:sampleBazel4 up-to-date:
  bazel-bin/sampleBazel4_/sampleBazel4
INFO: Elapsed time: 0.238s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Running command line: bazel-bin/sampleBazel4_/sampleBazel4 <args omitted>
{"level":"info","time":"2025-06-25 11:16:21.680","caller":"pkg/service.go:31","msg":"Starting server","port":"8080"}
```
