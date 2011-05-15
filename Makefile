#
# $Id: Makefile,v 5522914bff63 2011/05/15 02:22:12 pjardetzky $ 
#
# @author         $Author$
# @version        $Rev$
# @lastrevision   $Date$
#
include includes.mk

TOPDIR = $(shell pwd)
SUBDIRS = yui

all:
	@echo "making $@ in ${TOPDIR} ..."
	@${MakeSubdirs}

clean:
	@echo "making $@ in ${TOPDIR} ..."
	@${CleanRule}
	@${RM} ${APP_NAME}/build
	@${FIND} . -name *~ | xargs ${RM}
	@${MakeSubdirs}



