clean:
	bazel clean --expunge

generate:
	bazel run //:gazelle

build:
	bazel build //:sampleBazel3

run: clean build
	bazel run //:sampleBazel3

go-run:
	go run main.go
