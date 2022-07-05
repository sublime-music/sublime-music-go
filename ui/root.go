package ui

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/sumnerevans/sublime-music-next/adapters/base"
)

type RootWindow struct {
	*adw.ApplicationWindow

	MainStack      *adw.ViewStack
	PlayerControls *PlayerControls

	// Tabs
	AlbumsTab    *AlbumsTab
	ArtistsTab   *ArtistsTab
	BrowseTab    *BrowseTab
	PlaylistsTab *PlaylistsTab
}

func CreateRootWindow(app *adw.Application, currentAdapter base.Adapter) *RootWindow {
	window := &RootWindow{
		ApplicationWindow: adw.NewApplicationWindow(&app.Application),
	}

	window.SetTitle("Sublime Music")
	window.SetDefaultSize(1342, 756)

	builder := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/root.ui")
	rootBox := builder.GetObject("root").Cast().(*gtk.Box)
	window.SetContent(rootBox)

	window.MainStack = builder.GetObject("main-stack").Cast().(*adw.ViewStack)

	// Add in each of the tabs.
	window.AlbumsTab = CreateAlbumsTab()
	builder.GetObject("albums").Cast().(*gtk.Box).Append(window.AlbumsTab)

	window.ArtistsTab = CreateArtistsTab()
	builder.GetObject("artists").Cast().(*gtk.Box).Append(window.ArtistsTab)

	window.BrowseTab = CreateBrowseTab()
	builder.GetObject("browse").Cast().(*gtk.Box).Append(window.BrowseTab)

	window.PlaylistsTab = CreatePlaylistsTab(currentAdapter)
	builder.GetObject("playlists").Cast().(*gtk.Box).Append(window.PlaylistsTab)

	// Add in the player controls.
	window.PlayerControls = CreatePlayerControls()
	rootBox.InsertChildAfter(window.PlayerControls, window.MainStack)

	return window
}

func (rw *RootWindow) GoToPlaylist(playlistID *glib.Variant) {
	rw.MainStack.SetVisibleChildName("playlists")
}
