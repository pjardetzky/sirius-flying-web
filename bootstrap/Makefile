#
# Bootstrap
#
include ../includes.mk

PACKAGE=bootstrap
DIST=${PACKAGE}-3.3.6-dist
ZIPFILE=${DIST}.zip

APPDIR=../${APP_NAME}
CSSDIR=${APPDIR}/css
FONTDIR=${APPDIR}/fonts
JSDIR=${APPDIR}/js

all: lib install

lib: 
	@${UNZIP} ./${ZIPFILE}

install:
	@mkdir -p ${JSDIR}
	@mkdir -p ${CSSDIR}
	@mkdir -p ${FONTDIR}
	@${CP} ${DIST}/fonts/* ${FONTDIR}
	@${CP} ${DIST}/css/*min.css ${CSSDIR}
	@${CP} ${DIST}/js/*min.js ${JSDIR}

clean:
	@${CleanRule}
	@${RM} ${DIST}
	@${RM} ${FONTDIR}
	@${RM} ${CSSDIR}/bootstrap*.css
	@${RM} ${JSDIR}/bootstrap*.js


