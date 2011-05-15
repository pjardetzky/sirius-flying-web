#
# $Id: Makefile,v 5522914bff63 2011/05/15 02:22:12 pjardetzky $ 
#
# @author         $Author$
# @version        $Rev$
# @lastrevision   $Date$
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


