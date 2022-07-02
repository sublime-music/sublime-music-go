package main

import (
	"embed"

	log "github.com/sirupsen/logrus"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

//go:embed resources.gresource.xml
var gresourceFS embed.FS

func init() {
	data, err := gresourceFS.ReadFile("resources.gresource.xml")
	if err != nil {
		log.Fatal("Could not load resource")
		log.Fatal(err)
		return
	}
	log.Info(string(data))
	resource, err := gio.NewResourceFromData(glib.NewBytes(data))
	if err != nil {
		log.Fatal("Could not load resource")
		log.Fatal(err)
		return
	}
	gio.ResourcesRegister(resource)
}
