clean:
	bazel clean --expunge

generate:
	bazel run //:gazelle
	bazel mod tidy

# build前记得先去改一下api下面的BUILD文件
build-proto:
	bazel build //api/hello/v1:hello

# 移动pb/grpc文件到api下
move-proto:
	./copy_proto_outputs.sh

build:
	bazel build //:sampleBazel8

run:
	bazel run //:sampleBazel8 -- serve

go-run:
	go run main.go serve