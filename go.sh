#!/usr/bin/env bash
set -Eeu -o pipefail

script_path="$(cd "$(dirname "${0}")" || exit 1; pwd)"
cd "$script_path" || exit 1

#go_files=("${script_path}/main.go")

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
    echo "  fmt      run gofmt"
    exit 1
}

shift 1

#ldflags="$("${script_path}/getldflags.py")"
binary="${script_path}/dist/lico_$(uname -o | tr "[:upper:]" "[:lower:]")_$(uname -m)/lico"

case "${mode}" in
    "build")
        #go build -ldflags "$ldflags" -o "${script_path}/lico" -- "${go_files[@]}" "$@"
        goreleaser build --snapshot --clean
        mv "$binary" "${script_path}/lico"
        ;;
    "run")
        #go run -ldflags "$ldflags" -- "${go_files[@]}" "$@"
        "$script_path/$(basename "$0")" "build"
        "$script_path/lico" "$@"
        ;;
    "drun")
        "$script_path/$(basename "$0")" "run" -l "$script_path/lico.list" "$@"
        ;;
    "fmt")
        go mod tidy
        gofmt -l -w .
        ;;
    "")
        exit 1
        ;;
esac
