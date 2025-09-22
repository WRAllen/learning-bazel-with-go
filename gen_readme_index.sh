#!/bin/bash

README="README.md"
TMP_OUTPUT=".subprojects.tmp"

START_MARK="<!-- BEGIN SUBPROJECTS -->"
END_MARK="<!-- END SUBPROJECTS -->"

{
    echo "$START_MARK"
    echo "## 🧩 子项目索引"
    echo ""
    echo "| 子项目 | 简介 |"
    echo "|--------|------|"

    # 获取所有目录并按自定义顺序排序
    dirs=()
    for dir in */; do
        [[ "$dir" == .* ]] && continue
        [[ ! -d "$dir" ]] && continue
        dirs+=("$dir")
    done
    
    # 自定义排序：sampleBazel, sampleBazel2-9, sampleBazel10
    # 使用更精确的排序逻辑
    sorted_dirs=()
    
    # 首先添加sampleBazel（没有数字的）
    for dir in "${dirs[@]}"; do
        if [[ "$dir" == "sampleBazel/" ]]; then
            sorted_dirs+=("$dir")
        fi
    done
    
    # 然后按数字顺序添加sampleBazel2-10
    for i in {2..10}; do
        for dir in "${dirs[@]}"; do
            if [[ "$dir" == "sampleBazel$i/" ]]; then
                sorted_dirs+=("$dir")
            fi
        done
    done
    
    for dir in "${sorted_dirs[@]}"; do
        name="${dir%/}"
        readme="$dir/README.md"
        intro="*暂无介绍内容*"

        if [[ -f "$readme" ]]; then
            desc=$(awk '
        BEGIN { printing=0 }
        /^# +介绍/ { printing=1; next }
        /^# +运行方式/ { printing=0 }
        printing && NF > 0 {
          line = $0
          gsub(/^[ \t]+/, "", line)
          gsub(/([\\|])/, "\\\\&", line)
          lines = lines line "<br>"
        }
        END { print lines }
      ' "$readme")

            if [[ -n "$desc" ]]; then
                intro="$desc"
            fi
        fi

        echo "| 🔹 [$name](./$name) | $intro |"
    done

    echo "$END_MARK"
} >"$TMP_OUTPUT"

# macOS 不生成备份文件的删除操作
sed -i '' "/$START_MARK/,/$END_MARK/d" "$README"

# 如果文件最后一行不是空行，则加一个换行
if [ -n "$(tail -c1 "$README")" ]; then
    echo "" >>"$README"
fi

# 追加新的内容
cat "$TMP_OUTPUT" >>"$README"

rm -f "$TMP_OUTPUT"

echo "✅ README.md 已更新为表格形式的子项目索引"
git add README.md
