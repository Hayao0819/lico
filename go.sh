#!/usr/bin/env bash
set -Eeu -o pipefail

script_path="$(cd "$(dirname "${0}")" || exit 1; pwd)"
cd "$script_path" || exit 1

go_files=("${script_path}/main.go")
mode="${1-""}"

[[ -n "$mode" ]] || {
    echo "Debug tool for lico"
    echo
    echo "Usage: $0 [mode] [lico-args]"
    echo
    echo "Mode:"
    echo "  build    make executable file"
    echo "  run      run lico"
    echo "  drun     run lico with '-l $script_path/lico.list'"
    exit 1
}

shift 1

case "${mode}" in
    "build")
        go build -o "${script_path}/lico" -- "${go_files[@]}" "$@"
        ;;
    "run")
        go run -- "${go_files[@]}" "$@"
        ;;
    "drun")
        "$script_path/$(basename "$0")" "run" -l "$script_path/lico.list" "$@"
        ;;
    "")
        exit 1
        ;;
esac
