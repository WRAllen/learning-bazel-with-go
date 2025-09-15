
# 一次性生成 proto 和 gRPC 代码，包括 go_library
proto:
	bazel build //api/hello/v1:hello

# 运行 gazelle 更新 BUILD 文件
gazelle:
	bazel run //:gazelle

# 运行 gazelle 更新依赖
gazelle-update-repos:
	bazel run //:gazelle -- update-repos -from_file=go.mod

move:
	./copy_proto_outputs.sh

# 清理构建缓存
clean:
	bazel clean
