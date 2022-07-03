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

	playlistScrollWindow := builder.GetObject("playlist-details-content").Cast().(*gtk.ScrolledWindow)
	playlistDetailsFlap := builder.GetObject("playlist-details-flap").Cast().(*adw.Flap)
	vadjustment := playlistScrollWindow.VAdjustment()
	vadjustment.ConnectValueChanged(func() {
		playlistDetailsFlap.SetRevealFlap(vadjustment.Value() > 70)
	})

	return &PlaylistsTab{
		Leaflet: builder.GetObject("root").Cast().(*adw.Leaflet),
	}
}
