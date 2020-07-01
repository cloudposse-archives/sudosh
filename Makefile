COPYRIGHT_SOFTWARE:=Sudo Shell
COPYRIGHT_SOFTWARE_DESCRIPTION:=Sudo Shell provides a login shell that can be used to audit sessions

PATH:=$(PATH):$(GOPATH)/bin

include $(shell curl --silent -o .build-harness "https://raw.githubusercontent.com/cloudposse/build-harness/master/templates/Makefile.build-harness"; echo .build-harness)

build: go/build

install:
	install -m 0755 release/sudosh /usr/local/bin/sudosh
