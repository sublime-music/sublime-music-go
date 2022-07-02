package ui

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

type PlayerControls struct {
	*gtk.ActionBar
}

func CreatePlayerControls() *PlayerControls {
	builder := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/playercontrols.ui")
	pc := &PlayerControls{
		ActionBar: builder.GetObject("root").Cast().(*gtk.ActionBar),
	}

	return pc
}
