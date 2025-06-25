clean:
	bazel clean --expunge

generate:
	bazel run //:gazelle
	bazel mod tidy

build:
	bazel build //:sampleBazel4

run: clean generate build
	bazel run //:sampleBazel4 -- serve

go-run:
	go run main.go serve
