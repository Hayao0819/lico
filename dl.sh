#!/bin/sh

set -e -u

binary_arch="none"
binary_os="none"
tmpdir="${TMPDIR-"/tmp/"}/lico"
url=""
dest="/usr/local/bin/lico"

# Get os info
uname_march="$(uname -m | tr '[:upper:]' '[:lower:]')"
uname_os="$(uname -o | tr '[:upper:]' '[:lower:]')"

# Check and format os
case "${uname_march}" in
    "arm64" | "x86_64")
        binary_arch="$uname_march"
        ;;
esac

case "${uname_os}" in
    "gnu/linux")
        binary_os="Linux"
        ;;
    "darwin")
        binary_os="Darwin"
        ;;
    *"windows"*)
        echo "Windows is not supported by this script" >&2
        exit 1
    ;;
esac

if [ "$binary_arch" = "none" ] || [ "$binary_os" = "none" ]; then
    {
        echo "Unknown OS or Architecture"
        echo "OS: ${uname_os} Arch: ${uname_march}"
    } >&2
    exit 1
fi

# Check command line tools
for t in "curl" "tar"; do
    if ! which "$t" >/dev/null 2>&1; then
        echo "$t should be installed" >&2
        exit 1
    fi
done


# Run download
url="https://github.com/Hayao0819/lico/releases/latest/download/lico_${binary_os}_${binary_arch}.tar.gz"
echo "Downloading $url" >&2
mkdir -p "$tmpdir"
curl -f -L -o "$tmpdir/archive.tar.gz" "$url"

# Extract and install
tar xf "$tmpdir/archive.tar.gz" -C "$tmpdir"
sudo cp "$tmpdir/lico" "$dest"
echo "lico has been installed in $dest" >&2
