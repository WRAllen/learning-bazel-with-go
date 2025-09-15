
# 一次性生成 proto 和 gRPC 代码，包括 go_library
proto:
	bazel build //api/hello/v1:hello

move:
	./copy_proto_outputs.sh

# 清理构建缓存
clean:
	bazel clean
