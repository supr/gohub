include $(GOROOT)/src/Make.inc

TARG=gohub
GOFILES=\
		gohub.go\
		user.go\

include $(GOROOT)/src/Make.pkg

GOFMT=gofmt
BADFMT=$(shell $(GOFMT) -l $(GOFILES) 2> /dev/null)

gofmt: $(BADFMT)
			@for F in $(BADFMT); do $(GOFMT) -w $$F && echo $$F; done

ifneq ($(BADFMT),)
ifneq ($(MAKECMDGOALS), gofmt)
#$(warning WARNING: make gofmt: $(BADFMT))
endif
endif
