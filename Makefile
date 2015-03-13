TEST_MODULE = conf net humanize load io display
DESTDIR = 
bindir = /usr/bin
.PHONY: default clean install build test

default: build

clean:
	@rm -v cli/cli 2>/dev/null || :

install:
	mkdir -p $(DESTDIR)$(bindir)
	cp -a cli/cli $(DESTDIR)$(bindir)/tmux-top

build:
	cd cli ; go build

test:
	@for package in $(TEST_MODULE) ; do \
		cd $${package} && go test ; cd "../" ;\
	done
