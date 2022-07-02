package ui

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type ArtistsTab struct {
	*adw.Leaflet
}

func CreateArtistsTab() *ArtistsTab {
	builder := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/artists.ui")
	artistsTab := &ArtistsTab{
		Leaflet: builder.GetObject("root").Cast().(*adw.Leaflet),
	}

	return artistsTab
}
