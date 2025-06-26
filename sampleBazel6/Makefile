clean:
	bazel clean --expunge

generate:
	bazel run //:gazelle
	bazel mod tidy

build:
	bazel build //:sampleBazel5

run: clean generate build
	bazel run //:sampleBazel5 -- serve

only-run:
	bazel run //:sampleBazel5 -- serve

go-run:
	go run main.go serve
