
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate-proto:
	protoc \
		--go_out=. --go-grpc_out=. \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		api/*.proto

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
