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
    echo "  build             make executable file"
    echo "  install [Dir]     install lico to /usr/local/bin/ or ~/.bin/"
    echo "  run               run lico"
    echo "  drun | example    run lico with example config"
    echo "  fmt               run gofmt"
    echo "  newcmd NAME       make new command from template"
    echo "  release           make release"
    exit 1
}

shift 1

call_myself(){
    "${script_path}/$(basename "$0")" "$@" || return "$?"
}

run_build(){
    logfile="$(mktemp)"
    check_cmd "goreleaser"  || {
        echo "Please run this: go install github.com/goreleaser/goreleaser@latest"
        return 1
    }
    # 色付きでログを出力するにはもう少し工夫が必要
    # teeを使えばできるかもしれないが、pipefailをPOSIX Shellでは使えない
    goreleaser build --snapshot --clean --single-target 1> "$logfile" 2>&1 || {
        cat "$logfile"
        exit 1
    }
    rm -f "$logfile"

}

get_built_binary(){
    
    if [ ! -e "${script_path}/dist/artifacts.json" ]; then
        echo "Run 'run_build' before calling 'get_built_binary'" >&2
        return 1
    fi
    #"${script_path}/getpath.py"

    go run -- "${script_path}/tools/main.go" artifact "${script_path}/dist/artifacts.json"
}

check_cmd(){
    [ -n "${1-""}" ] || return 1
    if ! which "$1" >/dev/null 2>&1; then
        echo "$1 command is not installed." >&2
        return 1
    fi
    return 0
}

install_to(){
    mkdir -p "$1"
    run_build
    if ! echo "${PATH-""}" | tr ":" "\n" | grep -q "$(cd "$1"; pwd)"; then
        echo "Please add path to $1" >&2
    fi
    cp "$(get_built_binary)" "$1"
}

check_cmd go

case "${mode}" in
    "build")
        run_build
        mv "$(get_built_binary)" "${script_path}/lico"
        ;;
    "install")
        if [ -n "${1-""}" ]; then
            if [ ! -d "$1" ]; then
                echo "Please specify directory"
                exit 1
            else
                install_to "$1"
            fi
        elif [ "$(id -u)" = 0 ]; then
            install_to /usr/local/bin/
        elif [ -e "$HOME/.bin" ]; then
            install_to "$HOME/.bin"
        else
            echo "You should run this script as root to install lico" >&2
            exit 1
        fi
        ;;
    "run")
        run_build
        "$(get_built_binary)" "$@"
        ;;
    "grun")
        run_build
        sudo "$(get_built_binary)" -g "$@"
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
