#!/usr/bin/env bash
set -euo pipefail

DEFAULT_VERSION="1.23.0"
GO_VERSION="${1:-${GO_VERSION:-$DEFAULT_VERSION}}"
OS="${GO_OS:-linux}"
ARCH="${GO_ARCH:-amd64}"
INSTALL_DIR="${GO_INSTALL_DIR:-/usr/local}"
TMP_DIR="${TMPDIR:-/tmp}"

uname_s=$(uname -s | tr '[:upper:]' '[:lower:]')
uname_m=$(uname -m)

if [[ "${OS}" == "" ]]; then
  OS="$uname_s"
fi

if [[ "${ARCH}" == "" ]]; then
  ARCH="$uname_m"
fi

case "$uname_s" in
  darwin) OS="darwin" ;;
  linux) OS="linux" ;;
  *) echo "Unsupported OS: $uname_s" >&2; exit 1 ;;
esac

case "$uname_m" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $uname_m" >&2; exit 1 ;;
esac

if command -v go >/dev/null 2>&1; then
  current=$(go version | awk '{print $3}')
  if [[ "$current" == "go${GO_VERSION}" ]]; then
    echo "Go ${GO_VERSION} already installed."
    exit 0
  else
    echo "Found ${current}, will install go${GO_VERSION}."
  fi
fi

TARBALL="go${GO_VERSION}.${OS}-${ARCH}.tar.gz"
URL="https://go.dev/dl/${TARBALL}"
DEST="${TMP_DIR}/${TARBALL}"

echo "Downloading ${URL}..."
if command -v curl >/dev/null 2>&1; then
  curl -fsSL "$URL" -o "$DEST"
elif command -v wget >/dev/null 2>&1; then
  wget -q "$URL" -O "$DEST"
else
  echo "Neither curl nor wget is available" >&2
  exit 1
fi

if [[ ! -w "${INSTALL_DIR}" ]]; then
  SUDO=${SUDO:-sudo}
else
  SUDO=""
fi

if [[ -n "${SUDO}" ]]; then
  echo "Installing to ${INSTALL_DIR} with ${SUDO}."
else
  echo "Installing to ${INSTALL_DIR}."
fi

${SUDO} rm -rf "${INSTALL_DIR}/go"
${SUDO} tar -C "${INSTALL_DIR}" -xzf "$DEST"

cat <<MSG
Go ${GO_VERSION} installed to ${INSTALL_DIR}/go.
Add the following to your shell profile if it is not already present:

export GOROOT=${INSTALL_DIR}/go
export GOPATH=\${HOME}/go
export PATH=\${PATH}:\${GOROOT}/bin:\${GOPATH}/bin

MSG
