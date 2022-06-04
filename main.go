package main

import (
	"os"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("app.sublimemusic.SublimeMusic", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

type SublimeMusic struct {
	*gtk.Application

	window *gtk.ApplicationWindow
}

func activate(app *gtk.Application) {
	window := gtk.NewApplicationWindow(app)
	window.SetTitle("Sublime Music")
	window.SetTitlebar(createHeaderBar())

	window.SetChild(gtk.NewLabel("ohea"))
	audio := gtk.NewMediaFileForFilename("/home/sumner/tmp/ohea.mp3")
	audio.Play()

	window.SetDefaultSize(1342, 756)
	window.Show()

}

func createHeaderBar() *gtk.HeaderBar {
	headerBar := gtk.NewHeaderBar()
	headerBar.PackStart(gtk.NewLabel("Search here"))
	headerBar.SetTitleWidget(gtk.NewLabel("Tabs here"))
	headerBar.PackEnd(gtk.NewLabel("Settings and such here"))
	return headerBar
}
