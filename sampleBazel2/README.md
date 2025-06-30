# 介绍

这是一个多层次的简单go服务

主要用于展示go里面多个包之间存着引用时，BUILD要怎么写。

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
├── main.go
└── pkg
    ├── BUILD.bazel
    ├── controller.go
    └── log
        ├── BUILD.bazel
        └── log.go

3 directories, 11 files
```

可以看到每个目录下，都有一个BUILD文件，根目录，pkg，以及pkg/log这3个目录都有自己的BUILD.bazel

和sampleBazel发现多了一个deps的字段，里面声明了具体的依赖，bazel需要用户显示的在deps里面申明依赖（但是第三方的包比如go.mod里面的他就会自动的加载到deps，这个后续介绍）

也就是你用了xxx文件夹下的x，你就要在deps里面加上，不然build的时候就会报错了。

这里还用到了BUILD里面的import字段，在 Bazel 的 BUILD 文件中，importpath 字段必须与你在 Go 代码中实际使用的 import 路径保持一致

## 运行结果

运行输出如下

```
-> % make run  
bazel clean --expunge
INFO: Starting clean (this may take a while). Use --async if the clean takes more than several minutes.
bazel build //:sampleBazel2
Starting local Bazel server (8.3.0) and connecting to it...
INFO: Analyzed target //:sampleBazel2 (88 packages loaded, 5999 targets configured).
INFO: Found 1 target...
Target //:sampleBazel2 up-to-date:
  bazel-bin/sampleBazel2_/sampleBazel2
INFO: Elapsed time: 61.270s, Critical Path: 38.87s
INFO: 12 processes: 6 internal, 6 darwin-sandbox.
INFO: Build completed successfully, 12 total actions
bazel run //:sampleBazel2
INFO: Analyzed target //:sampleBazel2 (32 packages loaded, 316 targets configured).
INFO: Found 1 target...
Target //:sampleBazel2 up-to-date:
  bazel-bin/sampleBazel2_/sampleBazel2
INFO: Elapsed time: 0.690s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Running command line: bazel-bin/sampleBazel2_/sampleBazel2
🚀 Server is running on http://localhost:8080
```

可以发现http服务正常启动了，下面测试一下

```
-> % curl http://localhost:8080
Hello, World! 🌍

```

发现也是正常输出了对应的结果了。
