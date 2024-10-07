#!/bin/bash

PROJECT=nmt
# Use your desired buf version
BUF_VERSION=1.44.0
# buf is installed to ~/bin/your-project-name.
BIN_DIR=$HOME/bin/$PROJECT

curl -sSL \
	"https://github.com/bufbuild/buf/releases/download/v$BUF_VERSION/buf-$(uname -s)-$(uname -m)" \
	-o "$BIN_DIR/buf"
chmod +x "$BIN_DIR/buf"
