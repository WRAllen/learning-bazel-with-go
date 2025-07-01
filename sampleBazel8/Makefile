clean:
	bazel clean --expunge

generate:
	bazel run //:gazelle
	bazel mod tidy

# build前记得先去改一下api下面的BUILD文件
build-proto:
	bazel build //api/hello/v1:hello

move-proto:
	./copy_proto_outputs.sh

build:
	bazel build //:sampleBazel8

only-run:
	bazel run //:sampleBazel8 -- serve

go-run:
	go run main.go serve