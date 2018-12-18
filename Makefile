COPYRIGHT_SOFTWARE:=Sudo Shell
COPYRIGHT_SOFTWARE_DESCRIPTION:=Sudo Shell provides a login shell that can be used to audit sessions

PATH:=$(PATH):$(GOPATH)/bin

-include $(shell curl -sSL -o .build-harness "https://git.io/build-harness"; echo .build-harness)

setup:
	make init go:deps-build go:deps-dev go:deps go:lint 

build:
	make go:build

install:
	install -m 0755 release/sudosh /usr/local/bin/sudosh
