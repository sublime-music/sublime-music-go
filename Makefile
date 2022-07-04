CSS_FILES=$(wildcard ./resources/css/*.css)
ICON_FILES=$(wildcard ./resources/icons/scalable/actions/*.svg)
UI_FILES=$(wildcard ./resources/ui/*.ui)
RESOURCE_FILE_NAMES=css icons ui
RESOURCE_FILES=$(patsubst %,resources/%.gresource,$(RESOURCE_FILE_NAMES))

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

resources/css.gresource.xml: $(CSS_FILES)
	bash ./scripts/mkresource.sh $@ '' true $^

resources/icons.gresource.xml: $(ICON_FILES)
	bash ./scripts/mkresource.sh $@ \
		'preprocess="xml-stripblanks"' false $^

resources/ui.gresource.xml: $(UI_FILES)
	bash ./scripts/mkresource.sh $@ \
		'' false $^

.PHONY: all clean
