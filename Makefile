#
# $Id: Makefile,v 9959f289aa30 2011/09/25 05:16:53 pjardetzky $ 
#
include includes.mk

TOPDIR = $(shell pwd)
SUBDIRS = bootstrap

all:
	@echo "making $@ in ${TOPDIR} ..."
	@${MakeSubdirs}

clean:
	@echo "making $@ in ${TOPDIR} ..."
	@${CleanRule}
	@${FIND} . -name *~ | xargs ${RM}
	@${MakeSubdirs}



