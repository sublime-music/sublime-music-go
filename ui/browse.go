package ui

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type BrowseTab struct {
	*adw.Leaflet
}

func CreateBrowseTab() *BrowseTab {
	builder := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/browse.ui")
	browseTab := &BrowseTab{
		Leaflet: builder.GetObject("root").Cast().(*adw.Leaflet),
	}

	return browseTab
}
