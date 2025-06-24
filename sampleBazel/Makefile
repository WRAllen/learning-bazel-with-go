# 这里命令会清理sampleBazel下面Bazel自动生成的文件
clean:
	bazel clean --expunge

# 这里命令会去查找当前目录下所有可以打包的目标
# bazel build //:hello 执行工作区根目录下名为 hello 的 可执行目标（如 go_binary、py_binary 等）
# 当前项目里面只有一个hello，所以等价于这个
build:
	bazel build //...

# bazel run会去自动先build相关可执行文件，然后运行他
# 当前项目等价于 bazel run //:hello
run: clean build
	bazel run //...

# 正常直接使用go进行运行
go-run:
	go run main.go
