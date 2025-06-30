clean:
	bazel clean --expunge

generate:
	bazel run //:gazelle
	bazel mod tidy

build: generate
	bazel build //api/hello/v1:hello

