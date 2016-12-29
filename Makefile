.RECIPEPREFIX != ps

DESTDIR ?=
PREFIX  ?= /usr

all:
    go build

install:
    install -D -m 00755 haste-it $(DESTDIR)$(PREFIX)/bin/haste-it
    ln -s $(PREFIX)/bin/haste-it $(DESTDIR)$(PREFIX)/bin/haste

uninstall:
    unlink $(DESTDIR)$(PREFIX)/bin/haste
    rm $(DESTDIR)$(PREFIX)/bin/haste-it

clean:
    rm haste-it