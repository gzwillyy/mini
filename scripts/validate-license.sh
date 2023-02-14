#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

find_files() {
  find . -not \( \
    \( \
      -wholename './vendor' \
      -o -wholename '*testdata*' \
      -o -wholename '*third_party*' \
      -o -wholename '*grpc.pb.*' \
    \) -prune \
  \) \
  \( -name '*.go' -o -name '*.sh' \)
}

# 使用"|| :"忽略grep返回空时的错误码
failed_license_header=($(find_files | xargs grep -L 'Licensed under the Apache License, Version 2.0 (the "License")' || :))
if (( ${#failed_license_header[@]} > 0 )); then
  echo "某些源文件缺少许可证标头。"
  printf '%s\n' "${failed_license_header[@]}"
  exit 1
fi

# 使用"|| :"忽略grep返回空时的错误码
failed_copyright_header=($(find_files | xargs grep -L 'Copyright 2023 The Mini Authors' || :))
if (( ${#failed_copyright_header[@]} > 0 )); then
  echo "某些源文件缺少版权标头。"
  printf '%s\n' "${failed_copyright_header[@]}"
  exit 1
fi
