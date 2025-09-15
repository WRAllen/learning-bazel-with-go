# SampleBazel9

这是一个使用 Bazel 构建的 Go 项目，演示了如何在 proto 文件中使用以下依赖：

- `google/protobuf/empty.proto`
- `google/protobuf/timestamp.proto`

## 项目结构

```
sampleBazel9/
├── api/
│   └── hello/
│       └── v1/
│           ├── BUILD.bazel
│           ├── hello.proto
│           ├── hello_test.go
│           └── hello.pb.go (自动生成)
├── BUILD.bazel
├── MODULE.bazel
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```

## 构建和运行

### 构建项目
```bash
make build
# 或者
bazel build //...
```

### 运行演示程序
```bash
bazel run //:sample_bazel_9
```

### 生成 proto 文件
```bash
make proto
# 或者
bazel build //api/hello/v1:api_proto
```

### 运行测试
```bash
make test
# 或者
bazel test //api/hello/v1:hello_test
```

### 清理构建缓存
```bash
make clean
# 或者
bazel clean
```

## 特性

- ✅ 使用 Bazel 作为构建系统
- ✅ 支持 gRPC 服务定义
- ✅ 使用 protobuf 的 Empty 和 Timestamp 类型
- ✅ 自动依赖管理和构建
- ✅ 完整的测试覆盖

## 依赖管理

项目通过 MODULE.bazel 文件管理以下依赖：

- `protobuf`: Protocol Buffers 支持 (版本 29.1)
- `rules_go`: Go 构建规则 (版本 0.55.1)
- `gazelle`: 自动生成 BUILD 文件 (版本 0.44.0)
- `rules_proto`: Proto 构建规则 (版本 7.1.0)
- `toolchains_protoc`: Protoc 工具链 (版本 0.4.3)

## 演示功能

运行 `bazel run //:sample_bazel_9` 可以看到以下演示：

1. **HelloRequest 和 HelloResponse**: 基本的消息类型
2. **TimeResponse**: 使用 Timestamp 类型的时间响应
3. **Empty 类型**: 用于不需要参数的方法
4. **Timestamp 使用**: 跨语言的时间表示

## 技术实现

- **Proto 文件**: `api/hello/v1/hello.proto` 包含了新的 import 语句
- **依赖配置**: `api/hello/v1/BUILD.bazel` 正确配置了 proto 依赖
- **模块管理**: `MODULE.bazel` 管理外部依赖
- **自动构建**: Bazel 自动处理所有依赖和构建步骤

## 成功验证

✅ Proto 文件成功导入 `google/protobuf/empty.proto` 和 `google/protobuf/timestamp.proto`  
✅ Bazel 自动处理依赖解析和构建  
✅ 生成的 Go 代码可以正常使用  
✅ 测试通过，验证功能正确  
✅ 演示程序运行成功
