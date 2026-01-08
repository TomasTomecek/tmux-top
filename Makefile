TEST_MODULE = conf net humanize load io display sens disk
PREFIX ?= ${DESTDIR}/usr
INSTALLDIR=${PREFIX}/bin
DESTDIR =
bindir = /usr/bin
.PHONY: default clean install build test

default: build

clean:
	@rm -v cli/cli 2>/dev/null || :

install:
	install -D -m 755 tmux-top $(INSTALLDIR)/tmux-top

build:
	go build -o tmux-top ./cmd/tmux-top

test:
	@for package in $(TEST_MODULE) ; do \
		cd $${package} && go test ; cd "../" ;\
	done
