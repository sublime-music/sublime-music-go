package ui

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

type AlbumsTab struct {
	*gtk.Box
}

func CreateAlbumsTab() *AlbumsTab {
	builder := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/albums.ui")
	albumsTab := AlbumsTab{
		Box: builder.GetObject("root").Cast().(*gtk.Box),
	}

	return &albumsTab
}
