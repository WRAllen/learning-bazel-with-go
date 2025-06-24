clean:
	bazel clean --expunge

# 这里build直接用根目录下的BUILD里面的name
build:
	bazel build //:sampleBazel3

# 同上，也可以直接bazel run //...
run: clean build
	bazel run //:sampleBazel3

go-run:
	go run main.go
