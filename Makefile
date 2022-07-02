ICON_FILES=$(wildcard ./resources/icons/scalable/actions/*.svg)
UI_FILES=$(wildcard ./resources/ui/*.ui)
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

resources/ui.gresource.xml: $(UI_FILES)
	bash ./scripts/mkresource.sh $@ $^

.PHONY: all clean
