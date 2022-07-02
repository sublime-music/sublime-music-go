package ui

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type RootWindow struct {
	*adw.ApplicationWindow

	MainStack *adw.ViewStack
}

func CreateRootWindow(app *gtk.Application) *RootWindow {
	window := RootWindow{
		ApplicationWindow: adw.NewApplicationWindow(app),
		MainStack:         createMainStack(),
	}

	window.SetTitle("Sublime Music Next")

	builder := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/root.ui")
	rootBox := builder.GetObject("root-box").Cast().(*gtk.Box)
	window.SetContent(rootBox)

	// 	rootBox.Append(window.createHeaderBar())
	// 	rootBox.Append(window.MainStack)
	// 	rootBox.Append(CreatePlayerControls())
	// 	rootBox.Append(window.createFooterSwitcher())

	window.SetDefaultSize(1342, 756)
	return &window
}

func createMainStack() (vs *adw.ViewStack) {
	vs = adw.NewViewStack()
	vs.SetVExpand(true)

	albumsPage := vs.AddTitled(gtk.NewLabel("Albums"), "albums", "Albums")
	albumsPage.SetIconName("library-music-symbolic")

	artistsPage := vs.AddTitled(gtk.NewLabel("Artists"), "artists", "Artists")
	artistsPage.SetIconName("library-artists-symbolic")

	browsePage := vs.AddTitled(gtk.NewLabel("Browse"), "browse", "Browse")
	browsePage.SetIconName("columns-symbolic")

	playlistsPage := vs.AddTitled(gtk.NewLabel("Playlists"), "playlists", "Playlists")
	playlistsPage.SetIconName("playlist-symbolic")
	return
}
