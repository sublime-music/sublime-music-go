package ui

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func OpenPreferences(_ *glib.Variant) {
	builder := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/preferences.ui")
	builder.GetObject("root").Cast().(*adw.PreferencesWindow).Show()
}
