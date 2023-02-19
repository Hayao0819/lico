#!/bin/sh
set -eu #-o pipefail

script_path="$(cd "$(dirname "${0}")" || exit 1; pwd)"
cd "$script_path" || exit 1

#go_files=("${script_path}/main.go")

mode="${1-""}"

[ -n "$mode" ] || {
    echo "Debug tool for lico"
    echo
    echo "Usage: $0 [mode] [lico-args]"
    echo
    echo "Mode:"
    echo "  build        make executable file"
    echo "  run          run lico"
    echo "  drun         run lico with '-l $script_path/lico.list'"
    echo "  fmt          run gofmt"
    echo "  newcmd NAME  make new command from template"
    exit 1
}

shift 1

call_myself(){
    "${script_path}/$(basename "$0")" "$@"
}

build_cmd(){
    check_cmd "goreleaser"
    goreleaser build --snapshot --clean --single-target >&2
    "${script_path}/getpath.py" 
}

check_cmd(){
    [ -n "${1-""}" ] || return 1
    if ! which "$1" >/dev/null 2>&1; then
        echo "$1 command is not installed." >&2
        return 1
    fi
    return 0
}

check_cmd go

case "${mode}" in
    "build")
        mv "$(build_cmd)" "${script_path}/lico"
        ;;
    "run")
        #go run -ldflags "$ldflags" -- "${go_files[@]}" "$@"
        #"$script_path/$(basename "$0")" "build"
        "$(build_cmd)" "$@"
        ;;
    "drun")
        "$script_path/$(basename "$0")" "run" -l "$script_path/lico.list" "$@"
        ;;
    "fmt")
        check_cmd "gofmt"
        go mod tidy
        gofmt -l -w .
        ;;
    "newcmd")
        go run -- "$script_path/tools/main.go" newcmd "$script_path/misc/cmd.go.template" "${script_path}/cmd/${1}.go" "$1"
        ;;
    *)
        echo "No such command"
        call_myself ""
        exit 1
        ;;
esac
