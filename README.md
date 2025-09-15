# 项目介绍

Bazel主要是Google作为单仓（Monorepo）的支持，Bazel支持在一个多语言仓库里面进行统一的build，也就是一个仓库里面有python、golang、java等用他就可以进行很好的管理。

这是一个基于bazel的go的项目集合，主要目的在于学习如何使用bazel进行go项目的打包和运行，是一个自学项目，希望也能帮到其他正在入门学习Bazel的同学😄

# 依赖

使用brew安装Bazelisk（官方推荐的安装方式）

```shell
brew install bazelisk
```

如果是brew进行直接安装bazel，可能安装的版本比较老，这里我用bazelisk，安装好bazelisk其实你就安装上了最新版本的Bazel了，默认会在如下位置

```shell
-> % which bazel  
/usr/local/bin/bazel
```

本项目使用的是如下版本(后续教程版本>=8.2.1)：

```
Bazelisk version: 1.25.0
Build label: 8.2.1
Build target: @@//src/main/java/com/google/devtools/build/lib/bazel:BazelServer
Build time: Thu Apr 17 18:37:44 2025 (1744915064)
Build timestamp: 1744915064
Build timestamp as int: 1744915064
```

# 说明

为了避免创建多个Repo进行学习，这里统一放在一起，每个子目录下为具体的bazel项目，具体信息请根据子目录下的README.md进行了解

## 举例

本地测试时，可具体cd到子目录进行处理，下面用sampleBazel子目录进行演示通用流程

```
# 进入子目录
cd sampleBazel

# 阅读该目录的README.md文件

# 尝试直接运行服务
make run 

```

输出如下

```shell
-> % make run
bazel clean --expunge
INFO: Starting clean (this may take a while). Use --async if the clean takes more than several minutes.
bazel build //...
Starting local Bazel server (8.2.1) and connecting to it...
INFO: Analyzed target //:hello (83 packages loaded, 5981 targets configured).
INFO: Found 1 target...
Target //:hello up-to-date:
  bazel-bin/hello_/hello
INFO: Elapsed time: 62.661s, Critical Path: 37.68s
INFO: 10 processes: 6 internal, 4 darwin-sandbox.
INFO: Build completed successfully, 10 total actions
bazel run //...
INFO: Analyzed target //:hello (0 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //:hello up-to-date:
  bazel-bin/hello_/hello
INFO: Elapsed time: 0.230s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Running command line: bazel-bin/hello_/hello
Hello, Bazel Bzlmod!
```

发现成功输出了Hello的字符串😄

建议第一次接触Bazel的同学可以按照该项目的顺序进行阅读理解。

# 下面是具体的子项目和其介绍

<!-- BEGIN SUBPROJECTS -->
## 🧩 子项目索引

| 子项目 | 简介 |
|--------|------|
| 🔹 [sampleBazel](./sampleBazel) | 最简单的go结合Bazel的例子，go会在控制台输出一串字符串<br>这里主要介绍一个最简单的基于Bazel的go项目是怎么做的。<br> |
| 🔹 [sampleBazel2](./sampleBazel2) | 这是一个多层次的简单go服务<br>主要用于展示go里面多个包之间存着引用时，BUILD要怎么写。<br> |
| 🔹 [sampleBazel3](./sampleBazel3) | 这是一个多层次的有依赖第三方包的go的web项目<br>介绍了基于gazelle如何去自动更新BUILD.bazel文件<br> |
| 🔹 [sampleBazel4](./sampleBazel4) | 基于sampleBazel3，添加了cobra让项目变成一个cli项目，添加了新的文件夹<br>介绍了通过gazelle如何去自动生成BUILD.bazel文件，以及如何自动更新依赖<br> |
| 🔹 [sampleBazel5](./sampleBazel5) | 基于sampleBazel4，添加了配置文件相关的内容<br>介绍了在gazelle自动生成BUILD后，如何把配置文件一起打包<br> |
| 🔹 [sampleBazel6](./sampleBazel6) | 基于sampleBazel5，添加了proto相关的部分(暂时手动生成pb.go)<br>介绍了有proto文件时如何用bazel进行打包，启动一个grpc服务。<br> |
| 🔹 [sampleBazel7](./sampleBazel7) | 介绍通过bazel基于proto文件生成pb，grpc文件<br> |
| 🔹 [sampleBazel8](./sampleBazel8) | 基于sampleBazel7的自动生成pb/gprc，结合6的web项目<br>说明如何做到先写proto再写接口服务，然后build项目。<br> |
| 🔹 [sampleBazel9](./sampleBazel9) | 基于sampleBazel8的基础，介绍了proto里面如果有引用到第三方的依赖<br>要如何进行处理和编写<br> |
<!-- END SUBPROJECTS -->
