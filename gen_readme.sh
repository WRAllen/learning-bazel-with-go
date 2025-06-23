#!/bin/bash

README="README.md"
TMP_OUTPUT=".subprojects.tmp"

START_MARK="<!-- BEGIN SUBPROJECTS -->"
END_MARK="<!-- END SUBPROJECTS -->"

# 生成子项目说明段落
{
  echo "$START_MARK"
  echo "## 🧩 子项目索引"
  echo ""

  for dir in */; do
    [[ "$dir" == .* ]] && continue
    [[ ! -d "$dir" ]] && continue

    name="${dir%/}"
    readme="$dir/README.md"

    echo "### [$name](./$name)"

    if [[ -f "$readme" ]]; then
      awk '
        BEGIN { printing=0 }
        /^# +介绍/ { printing=1; next }
        /^# +运行方式/ { printing=0 }
        printing && NF > 0 { print "  " $0 }
      ' "$readme"
    else
      echo "  *暂无介绍*"
    fi

    echo ""
  done

  echo "$END_MARK"
} > "$TMP_OUTPUT"

# 如果 README 不含占位符，直接追加
if ! grep -q "$START_MARK" "$README"; then
  {
    cat "$README"
    echo ""
    cat "$TMP_OUTPUT"
  } > "${README}.new"
else
  # 替换标记之间内容
  awk -v start="$START_MARK" -v end="$END_MARK" -v tmpfile="$TMP_OUTPUT" '
    BEGIN { copying=1 }
    {
      if ($0 ~ start) {
        print
        while ((getline line < tmpfile) > 0) print line
        copying=0
      } else if ($0 ~ end) {
        copying=1
        next
      }
      if (copying) print
    }
  ' "$README" > "${README}.new"
fi

mv "${README}.new" "$README"
rm -f "$TMP_OUTPUT"

echo "✅ README.md 已更新子项目简介段"
