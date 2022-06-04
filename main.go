package main

import (
	"os"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := SublimeMusic{
		Application: adw.NewApplication("app.sublimemusic.SublimeMusic", gio.ApplicationFlagsNone),
	}
	app.ConnectActivate(func() { app.activate() })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

type SublimeMusic struct {
	*adw.Application

	window *adw.ApplicationWindow
}

func (s *SublimeMusic) activate() {
	s.window = adw.NewApplicationWindow(&s.Application.Application)
	s.window.SetTitle("Sublime Music")

	s.window.SetContent(gtk.NewLabel("ohea"))
	audio := gtk.NewMediaFileForFilename("/home/sumner/tmp/ohea.mp3")
	audio.Play()

	s.window.SetDefaultSize(1342, 756)
	s.window.Show()
}

func createHeaderBar() *adw.HeaderBar {
	headerBar := adw.NewHeaderBar()
	headerBar.PackStart(gtk.NewLabel("Search here"))
	headerBar.SetTitleWidget(gtk.NewLabel("Tabs here"))
	headerBar.PackEnd(gtk.NewLabel("Settings and such here"))
	return headerBar
}
