# 介绍

最简单的go结合Bazel的例子，go会在控制台输出一串字符串

这里主要介绍一个最简单的基于Bazel的go项目是怎么做的。

# 运行方式

```
# 清除Bazel生成的文件
make clean

# 使用Bazel进行打包
make build

# 使用Bazel运行程序，这里run已经做了上面的clean和build了
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
├── go.mod
└── main.go
```

除了go相关的文件，main和mod，还有项目相关的makefile，readme，其他的就是Bazel相关的文件

```
├── BUILD.bazel		手动编写，和具体文件相关
├── MODULE.bazel	手动编写，项目相关
├── MODULE.bazel.lock	这个文件是自动生成的，在build的时生成
```

`MODULE.bazel`(也可以不要后缀，直接写成BUILD，但是现在还是推荐带后缀，对于idea比较友好，并且官方说如果同时存在BUILD.bazel和BUILD优先加载前者) 是 Bazel 的 **Bzlmod 模块系统**的配置文件

`BUILD` 文件是 Bazel 的 **构建规则定义文件** ，控制项目的**源码如何被编译、打包、测试**等。（每个使用到的代码目录下都应该有一个这样的文件进行申明，具体后续项目介绍）

## 运行

通过如下命令进行打包

```
make build
```

可以看到build后其实在该目录下多了几个文件夹

```
├── bazel-bin -> /private/var/tmp/_bazel_xxxxx/eaa157a219e1f2ce63962859cc519431/execroot/_main/bazel-out/darwin_x86_64-fastbuild/bin
├── bazel-out -> /private/var/tmp/_bazel_xxxxx/eaa157a219e1f2ce63962859cc519431/execroot/_main/bazel-out
├── bazel-sampleBazel -> /private/var/tmp/_bazel_xxxxx/eaa157a219e1f2ce63962859cc519431/execroot/_main
├── bazel-testlogs -> /private/var/tmp/_bazel_xxxxx/eaa157a219e1f2ce63962859cc519431/execroot/_main/bazel-out/darwin_x86_64-fastbuild/testlogs
```

他其实是bazel生成文件的软连接，具体路径也在上面了

`bazel-bin`（构建产物输出目录）

* 所有 **构建目标（如 `go_binary`, `cc_binary`, `java_binary`）生成的可执行文件、静态资源、生成代码** 都放在这里。
* 是你平时运行 `bazel build //:target` 后最终产物的归宿。

`bazel-out`（所有中间构建产物）

* 包含 Bazel 的  **完整构建输出树** ，不仅是最终产物，还有所有中间产物、依赖库、分析数据等。
* 是 Bazel 的“工作目录”，供内部引用使用。
* 不建议直接修改或使用这个目录的文件。

`bazel-<workspace>`（工作目录的 execroot）

* 是 Bazel 执行构建/测试动作时的“执行根目录”（`execroot`）
* Bazel 会在这里创建 sandbox、模拟 workspace，所有构建动作在此目录中隔离运行。
* 某些构建调试中会看到“路径引用 execroot”，就是指这里。

`bazel-testlogs`（测试日志目录）

* 存放所有 `bazel test` 执行后的测试输出，包括：

  * `test.log`：测试控制台输出
  * `test.xml`：JUnit 风格报告
  * `status.txt`：测试结果状态（PASSED, FAILED 等）
* 是调试测试失败最重要的地方。

bazel run其实run的也是bin下面构建的产物
