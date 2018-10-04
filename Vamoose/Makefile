# vi: set ts=4 shiftwidth=4:

# Copyright 2018 Digital Aggregates Corporation
# Licensed under the terms in LICENSE.txt
# author:Chip Overclock
# mailto:coverclock@diag.com
# https://github.com/coverclock/com-diag-vamoose
# "Chip Overclock" is a registered trademark.
# "Digital Aggregates Corporation" is a registered trademark.

##
## This is a *highly* _experimental_ attempt to use the gccgo compiler. It's
## a lot more complicated than it really needs to be. In order to follow the
## typical directory layout of my projects, Go-based and otherwise, and to
## also support the directory layout expected of the Google go toolchain, a
## softlink from the GOPATH to the repo directory is used. This allows me to
## use the facilities provided by the go command, but also some tools I use
## to support my own workflow. Because the import paths in the Go source code
## assumes the GOPATH layout, this Makefile for the most part assumes that as
## well. It's not pretty.
##

.PHONY:	default

default:	all

########## Customizations

TITLE				:=	Vamoose

MAJOR				:=	5# API changes that may require that applications be modified.
MINOR				:=	0# Only functionality or features added with no legacy API changes.
BUILD				:=	0# Only bugs fixed with no API changes or new functionality.

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

########## Directories

BIN_DIR				:=	bin# Shell components
CMD_DIR				:=	cmd# Command components
OUT_DIR				:=	out# Build artifacts
PKG_DIR				:=	pkg# Package components
SRC_DIR				:=	src# Source directory
TMP_DIR				:=	/tmp

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

GOCCFLAGS			:=	-I $(OUT)
GOLDFLAGS			:=	-Wl,-R,$(GCCPATH)

GOSOURCE			:=	$(GOPATH)/$(SRC_DIR)# GOPATH/src
GOLINK				:=	$(GOSOURCE)/$(REPOSITORY)# GOPATH/src/github.com/coverclock/com-diag-vamoose
GOROOT				:=	$(GOLINK)/$(TITLE)# GOPATH/src/github.com/coverclock/com-diag-vamoose/Vamoose
GONAMESPACE			:=	$(REPOSITORY)/$(TITLE)# github.com/coverclock/com-diag-vamoose/Vamoose

########## Commands and Option Flags

GO					:=	go
GOCC				:=	$(CROSS_COMPILE)gccgo
AR					:=	$(CROSS_COMPILE)ar
RANLIB				:=	$(CROSS_COMPILE)ranlib
STRIP				:=	$(CROSS_COMPILE)strip

########## Build Artifacts

ABSOLUTEPACKAGES	:=	$(filter-out $(wildcard $(GOROOT)/$(PKG_DIR)/*/*_test.go), $(wildcard $(GOROOT)/$(PKG_DIR)/*/*.go))
RELATIVEPACKAGES	:=	$(subst $(GOROOT)/,,$(ABSOLUTEPACKAGES))

ABSOLUTECOMMANDS	:=	$(wildcard $(GOROOT)/$(CMD_DIR)/*/*.go)
RELATIVECOMMANDS	:=	$(subst $(GOROOT)/,,$(ABSOLUTECOMMANDS))

ABSOLUTESOURCES		:=	$(ABSOLUTEPACKAGES) $(ABSOLUTECOMMANDS)
RELATIVESOURCES		:=	$(RELATIVEPACKAGES) $(RELATIVECOMMANDS)

TARGETPACKAGES		:=	$(addprefix $(OUT)/$(GONAMESPACE)/,$(addsuffix .o,$(basename $(RELATIVEPACKAGES))))
TARGETCOMMANDS		:=	$(addprefix $(OUT)/$(BIN_DIR)/,$(notdir $(basename $(RELATIVECOMMANDS))))
TARGETGENERATED		:=	$(addprefix $(OUT)/$(BIN_DIR)/,$(GENERATED))
TARGETSCRIPTS		:=	$(addprefix $(OUT)/,$(basename $(wildcard $(BIN_DIR)/*.sh)))

TARGETSHARED		:=	$(OUT)/$(LIB_DIR)/$(PROJECT_SO).$(MAJOR).$(MINOR)
TARGETSHARED		+=	$(OUT)/$(LIB_DIR)/$(PROJECT_SO).$(MAJOR)
TARGETSHARED		+=	$(OUT)/$(LIB_DIR)/$(PROJECT_SO)

TARGETLIBRARIES		:=	$(TARGETPACKAGES)
TARGETPROGRAMS		:=	$(TARGETCOMMANDS) $(TARGETGENERATED) $(TARGETSCRIPTS)
TARGETALL			:=	$(TARGETLIBRARIES) $(TARGETPROGRAMS)

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

try:
	mkdir -p $(OUT)/$(GONAMESPACE)/$(PKG_DIR)
	$(GOCC) $(GOFLAGS) -c -o $(OUT)/$(GONAMESPACE)/$(PKG_DIR)/ticks.o $(PKG_DIR)/ticks/ticks.go $(GOLDFLAGS)
	$(GOCC) $(GOFLAGS) -c -o $(OUT)/$(GONAMESPACE)/$(PKG_DIR)/fletcher.o $(PKG_DIR)/fletcher/fletcher.go $(GOLDFLAGS)
	$(GOCC) $(GOFLAGS) -c -o $(OUT)/$(GONAMESPACE)/$(PKG_DIR)/throttle.o $(PKG_DIR)/throttle/throttle.go $(GOLDFLAGS)
	$(GOCC) $(GOFLAGS) -c -o $(OUT)/$(GONAMESPACE)/$(PKG_DIR)/harness.o $(PKG_DIR)/harness/harness.go $(GOLDFLAGS)
	$(GOCC) $(GOFLAGS) -c -o $(OUT)/$(GONAMESPACE)/$(PKG_DIR)/gcra.o $(PKG_DIR)/gcra/gcra.go $(GOLDFLAGS)
	$(GOCC) $(GOFLAGS) -c -o $(OUT)/$(GONAMESPACE)/$(PKG_DIR)/contract.o $(PKG_DIR)/contract/contract.go $(GOLDFLAGS)
	mkdir -p $(OUT)/$(GONAMESPACE)/$(CMD_DIR)
	$(GOCC) $(GOFLAGS)    -o $(OUT)/$(GONAMESPACE)/$(CMD_DIR)/fletch $(CMD_DIR)/fletch/fletch.go $(OUT)/$(GONAMESPACE)/$(PKG_DIR)/*.o $(GOLDFLAGS)
	$(GOCC) $(GOFLAGS)    -o $(OUT)/$(GONAMESPACE)/$(CMD_DIR)/shape $(CMD_DIR)/shape/shape.go $(OUT)/$(GONAMESPACE)/$(PKG_DIR)/*.o $(GOLDFLAGS)

########## Helpers

.PHONY:	depend manifest goroot

manifest:
	@for FF in $(ABSOLUTESOURCES); do echo -n "ABSOLUTESOURCE "; ls $$FF; done
	@for FF in $(RELATIVESOURCES); do echo -n "RELATIVESOURCE "; ls $$FF; done

artifacts:
	@for FF in $(TARGETPACKAGES);  do echo "TARGETPACKAGE   "$$FF; done
	@for FF in $(TARGETCOMMANDS);  do echo "TARGETCOMMAND   "$$FF; done
	@for FF in $(TARGETGENERATED); do echo "TARGETGENERATED "$$FF; done
	@for FF in $(TARGETSCRIPTS);   do echo "TARGETSCRIPT    "$$FF; done

goroot:
	test -n "$(GOPATH)"
	mkdir -p $(GOPATH)/$(SRC_DIR)/$(PROVIDER)/$(AUTHOR)
	cd $(GOPATH)/$(SRC_DIR)/$(PROVIDER)/$(AUTHOR); ln -f -s $(REPOROOT)

depend:	$(RELATIVEPACKAGES) $(RELATIVECOMMANDS)
	mkdir -p $(dir $@)
	cp /dev/null $@
	for FF in $(RELATIVESOURCES); do \
		echo "$(OUT)/$(GONAMESPACE)/$${FF%.go}.o:\t$$FF"; \
		PP=$$(${GO} list -f '{{ join .Deps "\n" }}' $$FF | grep '^$(GONAMESPACE)/' | tac); \
		for DD in $$PP; do \
			GG=$$(cd $(GOSOURCE); ls $$DD/*.go | grep -v "/[^/]*_test.go$$"); \
			for II in $$GG; do \
				echo "$(OUT)/$(GONAMESPACE)/$${FF%.go}.o:\t$(OUT)/$${II%.go}.o"; \
			done; \
		done; \
	done >> $@
	for FF in $(RELATIVECOMMANDS); do \
		DD=$$(dirname $$FF); \
		for II in $$DD; do \
			echo "$(OUT)/$(BIN_DIR)/$$(basename $$FF .go):\t$(OUT)/$${II%.go}.o"; \
		done; \
	done >> $@
	cat $@

-include $(OUT)/depend.mk

########## Target Go Packages

$(OUT)/$(ARC_DIR)/$(PROJECT_A):	$(TARGETOBJECTS)
	D=`dirname $@`; mkdir -p $$D
	$(AR) $(ARFLAGS) $@ $^
	$(RANLIB) $@

$(OUT)/$(LIB_DIR)/lib$(PROJECT).$(SO).$(MAJOR).$(MINOR):	$(TARGETOBJECTS)
	D=`dirname $@`; mkdir -p $$D
	$(CC) $(CFLAGS) $(A2SOARCH) -o $@ $(SOFLAGS) $(TARGETOBJECTS)

$(OUT)/$(LIB_DIR)/lib$(PROJECT).$(SO).$(MAJOR):	$(OUT)/$(LIB_DIR)/lib$(PROJECT).$(SO).$(MAJOR).$(MINOR)
	D=`dirname $<`; F=`basename $<`; T=`basename $@`; ( cd $$D; ln -s -f $$F $$T ) 

$(OUT)/$(LIB_DIR)/lib$(PROJECT).$(SO):	$(OUT)/$(LIB_DIR)/lib$(PROJECT).$(SO).$(MAJOR)
	D=`dirname $<`; F=`basename $<`; T=`basename $@`; ( cd $$D; ln -s -f $$F $$T ) 

########## Target Go Commands

$(OUT)/$(SYM_DIR)/%:	$(OUT)/$(OBC_DIR)/$(BIN_DIR)/%.o $(TARGETLIBRARIES)
	D=`dirname $@`; mkdir -p $$D
	$(CC) $(CFLAGS) -o $@ $< $(LDFLAGS)

$(OUT)/$(SYM_DIR)/%:	$(OUT)/$(OBX_DIR)/$(BIN_DIR)/%.o $(TARGETLIBRARIESXX) $(TARGETLIBRARIES)
	D=`dirname $@`; mkdir -p $$D
	$(CXX) $(CXXFLAGS) -o $@ $< $(LDXXFLAGS)

########## Generated Files

# For generating a setup script for a bash shell (for example, "bash generate > setup").
# (Because the Android bash doesn't seem to implement BASH_ARGV.)
$(OUT)/$(BIN_DIR)/generate:	Makefile
	D=`dirname $@`; mkdir -p $$D
	echo 'COM_DIAG_$(SYMBOL)_PATH=`dirname $$0`; COM_DIAG_$(SYMBOL)_ROOT=`cd $$COM_DIAG_$(SYMBOL)_PATH; pwd`' > $@
	echo 'echo export PATH=\\$$PATH:$$COM_DIAG_$(SYMBOL)_ROOT/../$(BIN_DIR):$$COM_DIAG_$(SYMBOL)_ROOT/../$(CMD_DIR)' >> $@
	echo 'echo export LD_LIBRARY_PATH=\\$$LD_LIBRARY_PATH:$$COM_DIAG_$(SYMBOL)_ROOT/../$(PKG_DIR):$(GCCPATH)' >> $@
	chmod 755 $@

# For sourcing into a bash shell (for example, ". setup").
$(OUT)/$(BIN_DIR)/setup:	$(OUT)/$(BIN_DIR)/generate
	$< > $@
	chmod 664 $@

########## Rules

$(OUT)/$(GONAMESPACE)/%.o:	%.go
	D=`dirname $@`; mkdir -p $$D
	$(GOCC) $(GOFLAGS) -o $@ -c $< $(GOLDFLAGS)
	

$(OUT)/$(GONAMESPACE)/%:	%.go
	D=`dirname $@`; mkdir -p $$D
	$(GOCC) $(GOFLAGS) -o $@ $< $(GOLDFLAGS)

$(OUT)/%:	%.sh
	D=`dirname $@`; mkdir -p $$D
	cp $< $@
	chmod 755 $@
