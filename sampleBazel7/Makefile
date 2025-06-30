clean:
	bazel clean --expunge

generate:
	bazel run //:gazelle
	bazel mod tidy

# build前记得先去改一下api下面的BUILD文件
build:
	bazel build //api/hello/v1:hello

