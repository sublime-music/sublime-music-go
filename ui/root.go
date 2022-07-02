package ui

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type RootWindow struct {
	*adw.ApplicationWindow

	MainStack      *adw.ViewStack
	PlayerControls *PlayerControls
}

func CreateRootWindow(app *gtk.Application) *RootWindow {
	window := RootWindow{
		ApplicationWindow: adw.NewApplicationWindow(app),
	}

	window.SetTitle("Sublime Music Next")
	window.SetDefaultSize(1342, 756)

	builder := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/root.ui")
	rootBox := builder.GetObject("root").Cast().(*gtk.Box)
	window.SetContent(rootBox)

	window.MainStack = builder.GetObject("main-stack").Cast().(*adw.ViewStack)

	// Add in each of the tabs.
	builder.GetObject("albums").Cast().(*gtk.Box).Append(CreateAlbumsTab())
	builder.GetObject("artists").Cast().(*gtk.Box).Append(CreateArtistsTab())
	builder.GetObject("browse").Cast().(*gtk.Box).Append(CreateBrowseTab())
	builder.GetObject("playlists").Cast().(*gtk.Box).Append(CreatePlaylistsTab())

	window.MainStack.SetVisibleChildName("playlists")

	// Add in the player controls.
	window.PlayerControls = CreatePlayerControls()
	rootBox.InsertChildAfter(window.PlayerControls, window.MainStack)

	return &window
}
