#
# Bootstrap
#
include ../includes.mk

PACKAGE=bootstrap
ZIPFILE=${PACKAGE}-3.0.3-dist.zip

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
	@${CP} dist/fonts/* ${FONTDIR}
	@${CP} dist/css/*min.css ${CSSDIR}
	@${CP} dist/js/*min.js ${JSDIR}

clean:
	@${CleanRule}
	@${RM} dist
	@${RM} ${FONTDIR}
	@${RM} ${CSSDIR}/bootstrap*.css
	@${RM} ${JSDIR}/bootstrap*.js


