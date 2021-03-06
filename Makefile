DIRS := \
	. \
	nl

DEPS = \
	github.com/vishvananda/netns

uniq = $(if $1,$(firstword $1) $(call uniq,$(filter-out $(firstword $1),$1)))
testdirs = $(call uniq,$(foreach d,$(1),$(dir $(wildcard $(d)/*_test.go))))
goroot = $(addprefix ../../../,$(1))
unroot = $(subst ../../../,,$(1))
fmt = $(addprefix fmt-,$(1))

all: test

$(call goroot,$(DEPS)):
	go get $(call unroot,$@)

.PHONY: $(call testdirs,$(DIRS))
$(call testdirs,$(DIRS)):
	sudo -E go test -test.parallel 4 -timeout 60s -v github.com/orkun1675/netlink/$@

$(call fmt,$(call testdirs,$(DIRS))):
	! gofmt -l $(subst fmt-,,$@)/*.go | grep -q .

.PHONY: fmt
fmt: $(call fmt,$(call testdirs,$(DIRS)))

test: fmt $(call goroot,$(DEPS)) $(call testdirs,$(DIRS))
