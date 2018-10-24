# vi: set ts=4 shiftwidth=4:

# Copyright 2018 Digital Aggregates Corporation
# Licensed under the terms in LICENSE.txt
# author:Chip Overclock
# mailto:coverclock@diag.com
# https://github.com/coverclock/com-diag-vamoose
# "Chip Overclock" is a registered trademark.
# "Digital Aggregates Corporation" is a registered trademark.

##
## This is a *highly* _experimental_ attempt to use the GNU gccgo compiler,
## an alternative to the Google gc compiler. This Makefile is a lot more
## complicated than it probably needs to be. In order to follow the canonical
## directory layout of my projects, Go-based and otherwise, and to also
## support the directory layout expected by the Google go toolchain, a
## softlink from the GOPATH to the repo directory is used. This allows me to
## use the facilities provided by the go command, but also some tools I use
## to support my own workflow. Because the import paths in the Go source code
## assumes the GOPATH layout, this Makefile has to deal both the repo layout
## and the GOPATH layout. It's not pretty.
##

.PHONY:	default

default:	all

########## Customizations

TITLE				:=	Vamoose

MAJOR				:=	7# API changes that may require that applications be modified.
MINOR				:=	0# Only functionality or features added with no legacy API changes.
BUILD				:=	1# Only bugs fixed with no API changes or new functionality.

PROJECT				:=	$(shell echo $(TITLE) | tr '[A-Z]' '[a-z]')
SYMBOL				:=	$(shell echo $(TITLE) | tr '[a-z]' '[A-Z]')

PROVIDER			:=	github.com
AUTHOR				:=	coverclock
REPOSITORY			:=	$(PROVIDER)/$(AUTHOR)/com-diag-$(PROJECT)
TARGET				:=	host

COPYRIGHT			:=	2018 Digital Aggregates Corporation
LICENSE				:=	GNU Lesser General Public License 2.1
CONTACT				:=	coverclock@diag.com
HOMEPAGE			:=	https://$(REPOSITORY)
HOST				:=	$(shell hostname -s)
BRANCH				:=	$(shell git rev-parse --abbrev-ref HEAD)
REVISION			:=	$(shell git rev-parse HEAD)
MODIFIED			:=	$(shell date -u -d @$(shell git log -1 --format="%at") +%Y-%m-%dT%H:%M:%S.%N%z)
ROOT				:=	$(shell git rev-parse --show-toplevel)
VINTAGE				:=	$(shell date -u +%Y-%m-%dT%H:%M:%S.%N%z)
OS					:=	$(shell uname -s)
KERNEL				:=	$(shell uname -r)
ARCH				:=	$(shell go version | cut -d ' ' -f 4)
PLATFORM			:=	$(shell go version | cut -d ' ' -f 3)

########## Directories

BIN_DIR				:=	bin# Shell components
CMD_DIR				:=	cmd# Command components
GEN_DIR				:=	gen# Generated sources
OUT_DIR				:=	out# Build artifacts
PKG_DIR				:=	pkg# Package components
SRC_DIR				:=	src# Source directory

TMPDIR				?=	/tmp
TMP_DIR				:=	$(TMPDIR)

########## Configuration

HERE				:=	$(shell pwd)

OUT					:=	$(OUT_DIR)/$(TARGET)

GITURL				:=	https://$(REPOSITORY).git

GENERATED			:=	generate setup

GCCPREFIX			:=	/usr
GCCDIRECTORY		:=	lib
GCCMACHINE			:=	x86_64-linux-gnu
GCCVERSION			:=	8
GCCPATH				:=	$(GCCPREFIX)/$(GCCDIRECTORY)/gcc/$(GCCMACHINE)/$(GCCVERSION)

GOSOURCE			:=	$(GOPATH)/$(SRC_DIR)# GOPATH/src
GOLINK				:=	$(GOSOURCE)/$(REPOSITORY)# GOPATH/src/github.com/coverclock/com-diag-vamoose
GOROOT				:=	$(GOLINK)/$(TITLE)# GOPATH/src/github.com/coverclock/com-diag-vamoose/Vamoose
GONAMESPACE			:=	$(REPOSITORY)/$(TITLE)# github.com/coverclock/com-diag-vamoose/Vamoose

CROSS_COMPILE		:=	$(TOOLCHAIN)

GOFLAGS				:=	-I $(OUT)
GOLDFLAGS			:=	-Wl,-R,$(GCCPATH)
GOLDMORE			:=	-L $(OUT)/$(GONAMESPACE)/$(PKG_DIR)
ARFLAGS				:=	crsv

DEPEND				:=	$(OUT)/depend.mk

########## Commands and Option Flags

GO					:=	go
GOCC				:=	$(CROSS_COMPILE)gccgo
GOLD				:=	$(CROSS_COMPILE)gccgo
AR					:=	$(CROSS_COMPILE)ar
RANLIB				:=	$(CROSS_COMPILE)ranlib
STRIP				:=	$(CROSS_COMPILE)strip

########## Build Sources, Modules, and Artifacts

ABSOLUTEPACKAGES	:=	$(wildcard $(GOROOT)/$(PKG_DIR)/*)
RELATIVEPACKAGES	:=	$(subst $(GOROOT)/,,$(ABSOLUTEPACKAGES))

ABSOLUTECOMMANDS	:=	$(wildcard $(GOROOT)/$(CMD_DIR)/*)
RELATIVECOMMANDS	:=	$(subst $(GOROOT)/,,$(ABSOLUTECOMMANDS))

ABSOLUTESOURCES		:=	$(filter-out $(wildcard $(GOROOT)/$(PKG_DIR)/*/*_test.go), $(wildcard $(GOROOT)/$(PKG_DIR)/*/*.go))
ABSOLUTESOURCES		+=	$(wildcard $(GOROOT)/$(CMD_DIR)/*/*.go)
RELATIVESOURCES		:=	$(subst $(GOROOT)/,,$(ABSOLUTESOURCES))

TARGETPACKAGES		:=	$(addprefix $(OUT)/$(GONAMESPACE)/,$(addsuffix .o,$(basename $(RELATIVEPACKAGES))))
TARGETPACKAGES		+=	$(OUT)/$(GONAMESPACE)/$(PKG_DIR)/$(PROJECT).o
TARGETCOMMANDS		:=	$(addprefix $(OUT)/$(BIN_DIR)/,$(notdir $(basename $(RELATIVECOMMANDS))))
TARGETGENERATED		:=	$(addprefix $(OUT)/$(BIN_DIR)/,$(GENERATED))
TARGETSCRIPTS		:=	$(addprefix $(OUT)/,$(basename $(wildcard $(BIN_DIR)/*.sh)))

TARGETPROGRAMS		:=	$(TARGETCOMMANDS) $(TARGETGENERATED) $(TARGETSCRIPTS)

TARGETALL			:=	$(TARGETPACKAGES) $(TARGETPROGRAMS)

########## Main Entry Points

.PHONY:	all tidy clean pristine clobber gopath

all:	$(TARGETALL)

tidy:
	rm -rf $(OUT)/$(BIN_DIR) $(OUT)/$(GONAMESPACE)/$(CMD_DIR) $(OUT)/$(GONAMESPACE)/$(PKG_DIR)

clean:
	rm -rf $(OUT)

pristine:	clean
	rm -rf $(OUT_DIR)

clobber:	pristine
	rm -f .cscope.lst .cscope.out .cscope.out.in .cscope.out.po

########## Helpers

.PHONY:	goroot list

goroot:
	test -n "$(GOPATH)"
	mkdir -p $(GOPATH)/$(SRC_DIR)/$(PROVIDER)/$(AUTHOR)
	cd $(GOPATH)/$(SRC_DIR)/$(PROVIDER)/$(AUTHOR); ln -f -s $(ROOT)
	rm -f $(DEPEND)

list:
	@for FF in $(ABSOLUTESOURCES); do echo -n "ABSOLUTESOURCE "; ls $$FF; done
	@for FF in $(RELATIVESOURCES); do echo -n "RELATIVESOURCE "; ls $$FF; done
	@for FF in $(ABSOLUTEPACKAGES); do echo -n "ABSOLUTEPACKAGE "; ls -d $$FF; done
	@for FF in $(RELATIVEPACKAGES); do echo -n "RELATIVEPACKAGE "; ls -d $$FF; done
	@for FF in $(ABSOLUTECOMMANDS); do echo -n "ABSOLUTECOMMAND "; ls -d $$FF; done
	@for FF in $(RELATIVECOMMANDS); do echo -n "RELATIVECOMMAND "; ls -d $$FF; done
	@for FF in $(TARGETPACKAGES);  do echo "TARGETPACKAGE   "$$FF; done
	@for FF in $(TARGETCOMMANDS);  do echo "TARGETCOMMAND   "$$FF; done
	@for FF in $(TARGETGENERATED); do echo "TARGETGENERATED "$$FF; done
	@for FF in $(TARGETSCRIPTS);   do echo "TARGETSCRIPT    "$$FF; done
	
########## Dependencies

.PHONY:	depend

$(DEPEND):	$(RELATIVESOURCES) $(RELATIVECOMMANDS) Makefile
	mkdir -p $(dir $@)
	echo "# GENERATED FILE: DO NOT EDIT; CHANGES WILL BE LOST!" > $@
	for FF in $(RELATIVESOURCES); do \
		MM=$$(dirname $$FF); \
		echo "$(OUT)/$(GONAMESPACE)/$$MM.o:\t$$FF"; \
		DD=$$(${GO} list -f '{{ join .Deps "\n" }}' $$FF | grep '^$(GONAMESPACE)/' | tac); \
		for PP in $$DD; do \
			echo "$(OUT)/$(GONAMESPACE)/$$MM.o:\t$(OUT)/$$PP.o"; \
		done; \
	done >> $@
	for FF in $(RELATIVECOMMANDS); do \
		echo "$(OUT)/$(BIN_DIR)/$$(basename $$FF):\t$(OUT)/$(GONAMESPACE)/$(CMD_DIR)/$$(basename $$FF).o"; \
	done >> $@

-include $(DEPEND)

depend:	$(DEPEND)
	cat $^

########## Generated Files

# For generating a setup script for a bash shell (for example, "bash generate > setup").
# (Because the Android bash doesn't seem to implement BASH_ARGV.)
$(OUT)/$(BIN_DIR)/generate:	Makefile
	DD=`dirname $@`; mkdir -p $$DD
	echo "# GENERATED FILE: DO NOT EDIT; CHANGES WILL BE LOST!" > $@
	echo 'COM_DIAG_$(SYMBOL)_PATH=`dirname $$0`; COM_DIAG_$(SYMBOL)_ROOT=`cd $$COM_DIAG_$(SYMBOL)_PATH; pwd`' >> $@
	echo 'echo export PATH=\\$$PATH:$$COM_DIAG_$(SYMBOL)_ROOT/../$(BIN_DIR):$$COM_DIAG_$(SYMBOL)_ROOT/../$(CMD_DIR)' >> $@
	echo 'echo export LD_LIBRARY_PATH=\\$$LD_LIBRARY_PATH:$$COM_DIAG_$(SYMBOL)_ROOT/../$(PKG_DIR):$(GCCPATH)' >> $@
	chmod 755 $@

# For sourcing into a bash shell (for example, ". setup").
$(OUT)/$(BIN_DIR)/setup:	$(OUT)/$(BIN_DIR)/generate
	$< > $@
	chmod 664 $@

$(OUT)/$(GEN_DIR)/$(PROJECT)/$(PROJECT).go:	Makefile
	D=`dirname $@`; mkdir -p $$D
	echo '/* GENERATED FILE: DO NOT EDIT; CHANGES WILL BE LOST! */' > $@
	echo 'package $(PROJECT)' >> $@
	echo 'const ARCH        string = "$(ARCH)"' >> $@
	echo 'const BRANCH      string = "$(BRANCH)"' >> $@
	echo 'const CONTACT     string = "$(CONTACT)"' >> $@
	echo 'const COPYRIGHT   string = "$(COPYRIGHT)"' >> $@
	echo 'const HOMEPAGE    string = "$(HOMEPAGE)"' >> $@
	echo 'const HOST        string = "$(HOST)"' >> $@
	echo 'const KERNEL      string = "$(KERNEL)"' >> $@
	echo 'const LICENSE     string = "$(LICENSE)"' >> $@
	echo 'const MODIFIED    string = "$(MODIFIED)"' >> $@
	echo 'const OS          string = "$(OS)"' >> $@
	echo 'const PLATFORM    string = "$(PLATFORM)"' >> $@
	echo 'const RELEASE     string = "$(MAJOR).$(MINOR).$(BUILD)"' >> $@
	echo 'const REPOSITORY  string = "$(REPOSITORY)"' >> $@
	echo 'const REVISION    string = "$(REVISION)"' >> $@
	echo 'const ROOT        string = "$(ROOT)"' >> $@
	echo 'const TARGET      string = "$(TARGET)"' >> $@
	echo 'const TITLE       string = "$(TITLE)"' >> $@
	echo 'const TOOLCHAIN   string = "$(TOOLCHAIN)"' >> $@
	echo 'const USER        string = "$(USER)"' >> $@
	echo 'const VINTAGE     string = "$(VINTAGE)"' >> $@

$(OUT)/$(GONAMESPACE)/$(PKG_DIR)/$(PROJECT).o:	$(OUT)/$(GEN_DIR)/$(PROJECT)/$(PROJECT).go
	DD=`dirname $@`; mkdir -p $$DD
	$(GOCC) $(GOFLAGS) -o $@ -c $(filter %.go,$^) $(GOLDFLAGS)

########## Rules

#	@echo "@="$@
#	@echo "^="$^
#	@echo "<="$<

$(OUT)/$(GONAMESPACE)/$(PKG_DIR)/%.o:	$(wildcard $(PKG_DIR)/%/*.go)
	DD=`dirname $@`; mkdir -p $$DD
	$(GOCC) $(GOFLAGS) -o $@ -c $(filter %.go,$^) $(GOLDFLAGS)

# The gold manual page indicates that the package object files must be
# "included in the link command". golang.org documentation about gccgo says
# the compiler does not "record the file name of imported packages in the
# object file. You must arrange for the imported data to be linked into the
# program." Empirical evidence supports this, unfortunately.

$(OUT)/$(GONAMESPACE)/$(CMD_DIR)/%.o:	$(wildcard $(CMD_DIR)/%/*.go)
	DD=`dirname $@`; mkdir -p $$DD
	$(GOCC) $(GOFLAGS) -o $@ -c $(filter %.go,$^) $(GOLDFLAGS)

$(OUT)/$(BIN_DIR)/%:	$(OUT)/$(GONAMESPACE)/$(CMD_DIR)/%.o
	DD=`dirname $@`; mkdir -p $$DD
	$(GOLD) $(GOFLAGS) -o $@ $< $(OUT)/$(GONAMESPACE)/$(PKG_DIR)/*.o $(GOLDFLAGS)

$(OUT)/%:	%.sh
	DD=`dirname $@`; mkdir -p $$DD
	cp $< $@
	chmod 755 $@
	
