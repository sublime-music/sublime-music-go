package ui

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type PlaylistsTab struct {
	*adw.Leaflet
}

func CreatePlaylistsTab() *PlaylistsTab {
	builder := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/playlists.ui")
	playlistsTab := &PlaylistsTab{
		Leaflet: builder.GetObject("root").Cast().(*adw.Leaflet),
	}

	return playlistsTab
}
