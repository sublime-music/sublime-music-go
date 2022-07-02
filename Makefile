ICON_FILES=$(wildcard ./resources/icons/scalable/actions/*.svg)
RESOURCE_FILES=resources/icons.gresource resources/ui.gresource

all: sublime-music-next

sublime-music-next: $(RESOURCE_FILES)
	go build -v .

run: $(RESOURCE_FILES)
	go run -v .

resources/%: resources/%.xml
	glib-compile-resources \
		--target=$@ \
		--sourcedir=resources \
		$<

resources/icons.gresource.xml: $(ICON_FILES)
	bash ./scripts/mkresource.sh $@ $^

resources/ui.gresource.xml: resources/ui/sublime-music.cmb
	cambalache --export-all $<
	bash ./scripts/mkresource.sh $@ resources/ui/*.ui

.PHONY: all clean
