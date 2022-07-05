package main

import (
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/sumnerevans/sublime-music-next/ui"
)

func addAction(app *SublimeMusic, name string, callback func(*glib.Variant), parameterType *glib.VariantType) {
	action := gio.NewSimpleAction(name, parameterType)
	action.ConnectActivate(callback)
	app.AddAction(action)
}

func (s *SublimeMusic) SetupActions() {
	addAction(s, "open-preferences", ui.OpenPreferences, nil)
	addAction(s, "go-to-playlist", s.window.GoToPlaylist, glib.NewVariantType("s"))

	// Add keyboard accelerators for some of the actions
	// TODO make this work
	s.SetAccelsForAction("open-preferences", []string{"<Ctrl>comma"})
}
