#!/usr/bin/env bash
set -euo pipefail

# === 可配置参数 ===
PROTO_PACKAGE="api/hello/v1"
BAZEL_TARGET="//${PROTO_PACKAGE}:hello"
TARGET_DIR="${PROTO_PACKAGE}"


# 构造输出目录路径（bazel-bin 下的路径）
OUTPUT_BASE="bazel-bin/${PROTO_PACKAGE}"

echo "🔍 Searching for .pb.go files under ${OUTPUT_BASE}..."
PB_FILES=$(find "${OUTPUT_BASE}" -type f \( -name "*_grpc.pb.go" -o -name "*.pb.go" \))

if [ -z "$PB_FILES" ]; then
    echo "❌ No generated .pb.go files found under ${OUTPUT_BASE}"
    exit 1
fi

# 确保目标目录存在
mkdir -p "${TARGET_DIR}"

echo "📄 Copying files to ${TARGET_DIR}:"
for f in $PB_FILES; do
    echo " → $(basename "$f")"
    cp -v "$f" "${TARGET_DIR}/"
done

echo "✅ Done."
