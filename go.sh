#!/usr/bin/env bash
set -Eeu -o pipefail

script_path="$(cd "$(dirname "${0}")" || exit 1; pwd)"
cd "$script_path" || exit 1

go_files=("${script_path}/main.go")

case "${1-""}" in
    "build")
        go build -o "${script_path}/lico" -- "${go_files[@]}"
        ;;
    "run")
        go run -- "${go_files[@]}"
        ;;
    "")
        exit 1
        ;;
esac
