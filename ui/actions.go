package ui

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

func addAction(app *adw.Application, name string, callback func(*glib.Variant), parameterType *glib.VariantType) {
	action := gio.NewSimpleAction(name, parameterType)
	action.ConnectActivate(callback)
	app.AddAction(action)
}

func SetupActions(app *adw.Application) {
	addAction(app, "open-preferences", OpenPreferences, nil)

	// Add keyboard accelerators for some of the actions
	// TODO make this work
	app.SetAccelsForAction("open-preferences", []string{"<Ctrl>comma"})
}
