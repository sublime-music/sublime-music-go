package ui

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

type PlayerControls struct {
	*gtk.ActionBar
}

func CreatePlayerControls() *PlayerControls {
	pc := PlayerControls{
		ActionBar: gtk.NewActionBar(),
	}

	pc.PackStart(gtk.NewLabel("Song info"))
	pc.PackEnd(gtk.NewLabel("Queue/volume info"))
	pc.SetCenterWidget(gtk.NewLabel("play button"))

	return &pc
}
