include ../../common.mk

SUBCOMMANDS = $(shell cd src && find subcommands -iname '*.go' | xargs -l dirname)

build-in-docker: clean
	docker run --rm \
		-v $$PWD/../..:$(GO_REPO_ROOT) \
		-w $(GO_REPO_ROOT)/plugins/repo \
		$(BUILD_IMAGE) \
		bash -c "make build" || exit $$?

build: commands subcommands
subcommands: $(SUBCOMMANDS)

commands: **/**/commands.go
	go build -a -o commands src/commands/commands.go

subcommands/%: src/subcommands/**/%.go
	go build -a -o $@ $<

clean:
	rm -rf commands subcommands
