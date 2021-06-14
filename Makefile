PKGNAME = haste-it
DESTDIR ?=
PREFIX  ?= /usr

all:
	CGO_ENABLED=0 go build -ldflags="-s -w"

validate:
	go vet ./...

install:
	install -D -m 00755 $(PKGNAME) $(DESTDIR)$(PREFIX)/bin/$(PKGNAME)
	ln -s $(PREFIX)/bin/$(PKGNAME) $(DESTDIR)$(PREFIX)/bin/haste

uninstall:
	unlink $(DESTDIR)$(PREFIX)/bin/haste
	rm $(DESTDIR)$(PREFIX)/bin/$(PKGNAME)

clean:
	rm $(PKGNAME)