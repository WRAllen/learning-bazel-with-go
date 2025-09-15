package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("SampleBazel9 - 演示新的 proto 功能")
	fmt.Println("=====================================")

	// 演示 HelloRequest 和 HelloResponse
	fmt.Println("\n1. 演示 HelloRequest 和 HelloResponse:")
	req := &api.HelloRequest{
		Name: "World",
	}

	resp := &api.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
		// CreatedAt 字段会在 proto 生成时自动包含
	}

	fmt.Printf("请求: %s\n", req.Name)
	fmt.Printf("响应: %s\n", resp.Message)
	if resp.CreatedAt != nil {
		fmt.Printf("创建时间: %s\n", resp.CreatedAt.AsTime().Format(time.RFC3339))
	}

	// 演示 TimeResponse
	fmt.Println("\n2. 演示 TimeResponse:")
	timeResp := &api.TimeResponse{
		Timezone: "Asia/Shanghai",
		// CurrentTime 字段会在 proto 生成时自动包含
	}

	if timeResp.CurrentTime != nil {
		fmt.Printf("当前时间: %s\n", timeResp.CurrentTime.AsTime().Format(time.RFC3339))
	}
	fmt.Printf("时区: %s\n", timeResp.Timezone)

	// 演示 Empty 类型的使用
	fmt.Println("\n3. 演示 Empty 类型:")
	fmt.Println("Empty 类型通常用于不需要参数的方法")
	fmt.Println("在我们的 proto 文件中，GetTime 方法使用了 Empty 作为输入参数")

	// 演示 Timestamp 的使用
	fmt.Println("\n4. 演示 Timestamp 的使用:")
	now := time.Now()
	fmt.Printf("当前时间: %s\n", now.Format(time.RFC3339))
	fmt.Println("Timestamp 类型提供了跨语言的时间表示")

	fmt.Println("\n✅ 所有功能演示完成！")
	fmt.Println("这个项目成功使用了以下新的 proto 功能:")
	fmt.Println("- google/protobuf/empty.proto")
	fmt.Println("- google/protobuf/timestamp.proto")
	fmt.Println("- 自动构建和依赖管理")
	fmt.Println("\n项目结构:")
	fmt.Println("- api/hello/v1/hello.proto: 包含新 import 的 proto 文件")
	fmt.Println("- api/hello/v1/BUILD.bazel: 配置了 proto 依赖")
	fmt.Println("- MODULE.bazel: 管理外部依赖")
}
