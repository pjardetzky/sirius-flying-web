#
# $Id: Makefile,v 9959f289aa30 2011/09/25 05:16:53 pjardetzky $ 
#
include ../includes.mk

PACKAGE=yui
ZIPFILE=${PACKAGE}_3.7.3.zip

YUIDIR=../${APP_NAME}
MENUDIR=${YUIDIR}/build/assets/node-menunav
MENUCSS=menu.css

all: lib menucss

lib: 
	@${UNZIP} ./${ZIPFILE}
	@${CP} yui/build ${YUIDIR}

menucss:
	@mkdir -p ${MENUDIR}
	@java -jar yuicompressor-2.4.8pre.jar ${MENUCSS} > ${MENUDIR}/${MENUCSS}

clean:
	@${RM} ${YUIDIR}/build
	@${CleanRule}
	@${RM} ${PACKAGE}


