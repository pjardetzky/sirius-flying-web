#
# $Id: includes.mk,v 9959f289aa30 2011/09/25 05:16:53 pjardetzky $ 
#
APP_NAME        = siriusflying

RM              = $(shell which rm) -rf
MV              = $(shell which mv)
CP              = $(shell which cp) -rp
UNZIP           = $(shell which unzip)
FIND            = $(shell which find)

#
# MakeSubdirs
#
# Typical usage:
#   SUBDIRS=test examples
#   @${MakeSubdirs}
#
# Arguments:
#   subdirs          - optional - specifies subdirs and overrides
#                                 SUBDIRS
#   target           - optional - specify the target to run. By default,
#                                 the current target is used.
#
define MakeSubdirs
(unset MAKEFLAGS;\
 set -e; \
 if [ "$$subdirs" = "" ]; then \
   subdirs='$(SUBDIRS)'; \
 fi; \
 if [ "$$target" = "" ]; then \
   target='$@'; \
 fi; \
 for dir in $$subdirs ; do \
   current=$(shell pwd); \
   if [ -f $$dir/Makefile ]; then \
   (cd $$dir ; echo "making $$target in $$current/$$dir..." ; \
    eval $(MAKE) $$k_make_flags $$target ) ; \
   if [ $$? != 0 ]; then \
     exit 1; \
   fi; \
   fi; \
 done; \
)
endef


#
# Rule for cleaning
#
define CleanRule
$(RM) *.o *.a *.so.* *~ *.d
endef


