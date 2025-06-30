clean:
	bazel clean --expunge

generate:
	bazel run //:gazelle
	bazel mod tidy

# 根据proto文件生成pb文件
build-pb: generate
	bazel build //api/hello/v1:api_go_proto

build: build-pb
	bazel build //:sampleBazel7

only-run:
	bazel run //:sampleBazel7 -- serve

go-run:
	go run main.go serve
