package resources

import (
	"embed"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	log "github.com/sirupsen/logrus"
)

//go:embed *.gresource
var gresourceFS embed.FS

func loadResource(resourceName string) error {
	data, err := gresourceFS.ReadFile(resourceName)
	if err != nil {
		log.Fatal("Failed to read resource")
		return err
	}
	resource, err := gio.NewResourceFromData(glib.NewBytes(data))
	if err != nil {
		return err
	}
	gio.ResourcesRegister(resource)
	return nil
}

func LoadResources() {
	// Load all of the resources
	entries, _ := gresourceFS.ReadDir(".")
	for _, entry := range entries {
		log.Debugf("Loading resource %s", entry.Name())
		err := loadResource(entry.Name())
		if err != nil {
			log.Fatalf("Failed to load resource %s", entry.Name())
		}
	}
}
