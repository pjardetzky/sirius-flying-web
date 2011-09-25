#
# $Id: Makefile,v 9959f289aa30 2011/09/25 05:16:53 pjardetzky $ 
#
include ../includes.mk

PACKAGE=yui
ZIPFILE=${PACKAGE}_2.7.0b.zip

all:
	@${UNZIP} ./${ZIPFILE}
	@cat reset-fonts-grids.patch >> yui/build/reset-fonts-grids/reset-fonts-grids.css
	@${CP} yui/build ../${APP_NAME}

clean:
	@${CleanRule}
	@${RM} ${PACKAGE}


