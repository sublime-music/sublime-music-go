package ui

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type PlayerControls struct {
	*adw.Squeezer
}

func CreatePlayerControls() *PlayerControls {
	builderBig := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/playercontrols.ui")
	builderSmall := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/playercontrolssmall.ui")

	squeezer := adw.NewSqueezer()
	squeezer.Add(builderBig.GetObject("root").Cast().(*gtk.ActionBar))
	squeezer.Add(builderSmall.GetObject("root").Cast().(*gtk.Box))
	squeezer.SetHomogeneous(false)
	
	pc := &PlayerControls{
		Squeezer: squeezer,
	}

	return pc
}
