#!/bin/bash

BUF_VERSION=1.44.0

curl -sSL \
	"https://github.com/bufbuild/buf/releases/download/v$BUF_VERSION/buf-$(uname -s)-$(uname -m)" \
	-o "./buf"
chmod +x "./buf"
