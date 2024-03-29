// Package resources loads all of the UI resources that GTK needs.
package resources

import (
	"embed"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/rs/zerolog/log"
)

//go:embed *.gresource
var gresourceFS embed.FS

func loadResource(resourceName string) error {
	data, err := gresourceFS.ReadFile(resourceName)
	if err != nil {
		log.Fatal().Msg("Failed to read resource")
		return err
	}
	resource, err := gio.NewResourceFromData(glib.NewBytes(data))
	if err != nil {
		return err
	}
	gio.ResourcesRegister(resource)
	return nil
}

// LoadResources loads all of the UI resources.
func LoadResources() {
	// Load all of the resources
	entries, _ := gresourceFS.ReadDir(".")
	for _, entry := range entries {
		log.Debug().Str("entry", entry.Name()).Msg("Loading resources")
		err := loadResource(entry.Name())
		if err != nil {
			log.Fatal().Str("entry", entry.Name()).Msg("Failed to load resource")
		}
	}
}
