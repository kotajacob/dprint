# dprint
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: clean build

build:
	go build
	scdoc < dprint.1.scd | sed "s/VERSION/$(VERSION)/g" > dprint.1

clean:
	rm -f dprint
	rm -f dprint.1

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f dprint $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/dprint
	mkdir -p $(DESTDIR)$(MANPREFIX)/man1
	cp -f dprint.1 $(DESTDIR)$(MANPREFIX)/man1/dprint.1
	chmod 644 $(DESTDIR)$(MANPREFIX)/man1/dprint.1

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/dprint
	rm -f $(DESTDIR)$(MANPREFIX)/man1/dprint.1

.PHONY: all build clean install uninstall
