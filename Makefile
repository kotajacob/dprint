# dprint
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: dprint dprint.1

dprint:
	$(GO) build $(GOFLAGS) \
		-ldflags "-X main.Version=$(VERSION) \
		-X main.Config=$(CONFIG)"

dprint.1:
	scdoc < dprint.1.scd | sed "s/VERSION/$(VERSION)/g" > dprint.1

clean:
	$(RM) dprint
	$(RM) dprint.1

install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f dprint $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/dprint
	mkdir -p $(DESTDIR)$(MANPREFIX)/man1
	cp -f dprint.1 $(DESTDIR)$(MANPREFIX)/man1/dprint.1
	chmod 644 $(DESTDIR)$(MANPREFIX)/man1/dprint.1

uninstall:
	$(RM) $(DESTDIR)$(PREFIX)/bin/dprint
	$(RM) $(DESTDIR)$(MANPREFIX)/man1/dprint.1

.DEFAULT_GOAL := all

.PHONY: all clean install uninstall
