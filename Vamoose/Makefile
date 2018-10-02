# vi: set ts=4 shiftwidth=4:
# Copyright 2018 Digital Aggregates Corporation
# Licensed under the terms in LICENSE.txt
# author:Chip Overclock
# mailto:coverclock@diag.com
# https://github.com/coverclock/com-diag-vamoose
# "Chip Overclock" is a registered trademark.
# "Digital Aggregates Corporation" is a registered trademark.

##
## This is a *highly* _experimental_ attempt to use the gccgo compiler.
##

.PHONY:	default

default:	try

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
RESOURCE			:=	$(REPOSITORY)/$(TITLE)
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

########## Configuration

HERE				:=	$(shell pwd)

OUT					:=	$(OUT_DIR)/$(TARGET)

TEMP_DIR			:=	/tmp

GITURL				:=	https://$(REPOSITORY).git

GENERATED			:=	generate setup

GCCGOPREFIX			:=	/usr
GCCGODIRECTORY		:=	lib
GCCGOMACHINE		:=	x86_64-linux-gnu
GCCGOVERSION		:=	8
GCCGOPATH			:=	$(GCCGOPREFIX)/$(GCCGODIRECTORY)/gcc/$(GCCGOMACHINE)/$(GCCGOVERSION)

GOFLAGS				:=	-I $(OUT)
GOLDFLAGS			:=	-Wl,-R,$(GCCGOPATH)

########## Commands and Option Flags

GO					:=	$(CROSS_COMPILE)gccgo
AR					:=	$(CROSS_COMPILE)ar
RANLIB				:=	$(CROSS_COMPILE)ranlib
STRIP				:=	$(CROSS_COMPILE)strip

########## Build Artifacts

TARGETPACKAGES		:=	$(addprefix $(OUT)/,$(addsuffix .o,$(basename $(filter-out $(GOPATH)/$(SRC_DIR)/$(RESOURCE)/$(PKG_DIR)/*/*_test.go,$(wildcard $(GOPATH)/$(SRC_DIR)/$(RESOURCE)/$(PKG_DIR)/*/*.go)))))
TARGETCOMMANDS		:=	$(addprefix $(OUT)/,$(basename $(wildcard $(GOPATH)/$(SRC_DIR)/$(RESOURCE)/$(CMD_DIR)/*/*.go)))
TARGETGENERATED		:=	$(addprefix $(OUT)/$(BIN_DIR)/,$(GENERATED)) $(addprefix $(OUT)/$(SYM_DIR)/,$(GENERATED))
TARGETSCRIPTS		:=	$(addprefix $(OUT)/,$(basename $(wildcard $(BIN_DIR)/*.sh)))

stuff:
	echo $(TARGETPACKAGES)
	echo $(TARGETCOMMANDS)

TARGETSHARED		:=	$(OUT)/$(LIB_DIR)/$(PROJECT_SO).$(MAJOR).$(MINOR)
TARGETSHARED		+=	$(OUT)/$(LIB_DIR)/$(PROJECT_SO).$(MAJOR)
TARGETSHARED		+=	$(OUT)/$(LIB_DIR)/$(PROJECT_SO)

TARGETLIBRARIES		:=	$(TARGETPACKAGES)
TARGETPROGRAMS		:=	$(TARGETCOMMANDS) $(TARGETGENERATED) $(TARGETSCRIPTS)
TARGETALL			:=	$(TARGETLIBRARIES) $(TARGETPROGRAMS)

########## Main Entry Points

.PHONY:	all tidy clean pristine clobber

all:	$(TARGETALL)

tidy:
	rm -rf $(OUT)/$(BIN_DIR) $(OUT)/$(CMD_DIR) $(OUT)/$(PKG_DIR)

clean:
	rm -rf $(OUT)

pristine:	clean
	rm -rf $(OUT_DIR)

clobber:	pristine
	rm -f .cscope.lst .cscope.out .cscope.out.in .cscope.out.po

try:
	mkdir -p $(OUT)/$(RESOURCE)/$(PKG_DIR)
	$(GO) $(GOFLAGS) -c -o $(OUT)/$(RESOURCE)/$(PKG_DIR)/ticks.o $(PKG_DIR)/ticks/ticks.go $(GOLDFLAGS)
	$(GO) $(GOFLAGS) -c -o $(OUT)/$(RESOURCE)/$(PKG_DIR)/fletcher.o $(PKG_DIR)/fletcher/fletcher.go $(GOLDFLAGS)
	$(GO) $(GOFLAGS) -c -o $(OUT)/$(RESOURCE)/$(PKG_DIR)/throttle.o $(PKG_DIR)/throttle/throttle.go $(GOLDFLAGS)
	$(GO) $(GOFLAGS) -c -o $(OUT)/$(RESOURCE)/$(PKG_DIR)/harness.o $(PKG_DIR)/harness/harness.go $(GOLDFLAGS)
	$(GO) $(GOFLAGS) -c -o $(OUT)/$(RESOURCE)/$(PKG_DIR)/gcra.o $(PKG_DIR)/gcra/gcra.go $(GOLDFLAGS)
	$(GO) $(GOFLAGS) -c -o $(OUT)/$(RESOURCE)/$(PKG_DIR)/contract.o $(PKG_DIR)/contract/contract.go $(GOLDFLAGS)
	mkdir -p $(OUT)/$(RESOURCE)/$(CMD_DIR)
	$(GO) $(GOFLAGS)    -o $(OUT)/$(RESOURCE)/$(CMD_DIR)/fletch $(CMD_DIR)/fletch/fletch.go $(OUT)/$(RESOURCE)/$(PKG_DIR)/*.o $(GOLDFLAGS)
	$(GO) $(GOFLAGS)    -o $(OUT)/$(RESOURCE)/$(CMD_DIR)/shape $(CMD_DIR)/shape/shape.go $(OUT)/$(RESOURCE)/$(PKG_DIR)/*.o $(GOLDFLAGS)

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

$(OUT)/%.o:	%.go
	D=`dirname $@`; mkdir -p $$D
	$(GO) $(GOFLAGS) -o $@ -c $<

$(OUT)/%:	%.go
	D=`dirname $@`; mkdir -p $$D
	$(GO) $(GOFLAGS) -o $@ -c $<

$(OUT)/%:	%.sh
	D=`dirname $@`; mkdir -p $$D
	cp $< $@
	chmod 755 $@
