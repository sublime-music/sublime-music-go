RESOURCE_XML_FILES=$(wildcard ./resources/*.gresource.xml)
RESOURCE_FILES=$(patsubst %.xml,%,$(RESOURCE_XML_FILES))

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

.PHONY: all clean
