#!/bin/sh
set -e -u #-o pipefail

script_path="$(cd "$(dirname "${0}")" || exit 1; pwd)"
cd "$script_path" || exit 1


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
    "${script_path}/$(basename "$0")" "$@" || return "$?"
}

run_build(){
    check_cmd "goreleaser"  || {
        echo "Please run this: go install github.com/goreleaser/goreleaser@latest"
        return 1
    }
    goreleaser build --debug --snapshot --clean --single-target >&2
}

get_built_binary(){
    if [ ! -e "${script_path}/dist/artifacts.json" ]; then
        echo "Run 'run_build' before calling 'get_built_binary'" >&2
        return 1
    fi
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
        run_build
        mv "$(get_built_binary)" "${script_path}/lico"
        ;;
    "install")
        call_myself "build"
        if [ "$(id -u)" = 0 ]; then
            cp "$script_path/lico" /usr/local/bin/
        elif [ -e "$HOME/.bin" ]; then
            cp "$script_path/lico" "$HOME/.bin"
        else
            echo "You should run this script as root to install lico" >&2
            exit 1
        fi
        ;;
    "run")
        run_build
        "$(get_built_binary)" "$@"
        ;;
    "drun" | "example")
        call_myself "run" -l "${script_path}/example/lico.list" -r "${script_path}/example" --created-list "$script_path/example/created.list" "$@"
        ;;
    "fmt")
        check_cmd "gofmt"
        go mod tidy
        gofmt -l -w .
        ;;
    "newcmd")
        go run -- "$script_path/tools/main.go" newcmd "$script_path/misc/cmd.go.template" "${script_path}/cmd/${1}.go" "$1"
        ;;
    "release")
        check_cmd "goreleaser"
        goreleaser release --snapshot --clean
        ;;
    *)
        echo "No such command" >&2
        call_myself ""
        exit 1
        ;;
esac
