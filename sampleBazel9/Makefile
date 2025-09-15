.PHONY: build test clean proto

# 构建项目
build:
	bazel build //...

# 运行测试
test:
	bazel test //...

# 清理构建缓存
clean:
	bazel clean

# 生成 proto 文件
proto:
	bazel build //api/hello/v1:api_proto

# 运行 gazelle 更新 BUILD 文件
gazelle:
	bazel run //:gazelle

# 更新依赖
update-deps:
	bazel run //:gazelle -- update-repos -from_file=go.mod
