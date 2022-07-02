package main

import (
	"os"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/sumnerevans/sublime-music-next/ui"
)

func main() {
	app := SublimeMusic{
		Application: adw.NewApplication("app.sublimemusic.SublimeMusicNext", gio.ApplicationFlagsNone),
	}
	app.ConnectActivate(func() { app.activate() })

	// Exit with whatever code the app exits with
	os.Exit(app.Run(os.Args))
}

type SublimeMusic struct {
	*adw.Application

	window *ui.RootWindow
}

func (s *SublimeMusic) activate() {
	s.window = ui.CreateRootWindow(&s.Application.Application)

	audio := gtk.NewMediaFileForFilename("/home/sumner/tmp/ohea.mp3")
	audio.Play()

	s.window.Show()
}
