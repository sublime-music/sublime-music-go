package ui

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"

	log "github.com/sirupsen/logrus"
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

	{
		rootBox := gtk.NewBox(gtk.OrientationVertical, 0)
		rootBox.SetVExpand(true)

		rootBox.Append(window.createHeaderBar())
		rootBox.Append(window.MainStack)
		rootBox.Append(CreatePlayerControls())

		window.SetContent(rootBox)
	}

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

func (w *RootWindow) createHeaderBar() *adw.HeaderBar {
	headerBar := adw.NewHeaderBar()

	// Search Button
	{
		menuButton := gtk.NewMenuButton()
		menuButton.SetIconName("loupe-symbolic")
		menuButton.Connect("activate", func(menuBtn *gtk.MenuButton) {
			log.Info("Activated")
		})
		headerBar.PackStart(menuButton)
	}

	// Stack Switcher
	{
		stackSwitcher := adw.NewViewSwitcherTitle()
		stackSwitcher.SetStack(w.MainStack)
		headerBar.SetTitleWidget(stackSwitcher)
	}

	headerBar.PackEnd(gtk.NewLabel("Settings and such here"))
	return headerBar
}
