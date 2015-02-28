TEST_MODULE = conf net humanize load io
.PHONY: default clean install build test

default: build

clean:
	@rm -v cli/cli 2>/dev/null || :

install:
	cp -a cli/cli /usr/bin/tmux-top

build:
	cd cli ; go build

test:
	@for package in $(TEST_MODULE) ; do \
		cd $${package} && go test ; cd "../" ;\
	done
