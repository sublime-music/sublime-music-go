package ui

import (
	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	log "github.com/sirupsen/logrus"

	"github.com/sumnerevans/sublime-music-next/adapters/base"
)

type PlaylistsTab struct {
	*adw.Leaflet

	currentAdapter base.Adapter
	playlistList   *gtk.ListView
}

func CreatePlaylistsTab(currentAdapter base.Adapter) *PlaylistsTab {
	builder := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/playlists.ui")

	playlistScrollWindow := builder.GetObject("playlist-details-content").Cast().(*gtk.ScrolledWindow)
	playlistDetailsFlap := builder.GetObject("playlist-details-flap").Cast().(*adw.Flap)
	vadjustment := playlistScrollWindow.VAdjustment()
	vadjustment.ConnectValueChanged(func() {
		playlistDetailsFlap.SetRevealFlap(vadjustment.Value() > 70)
	})

	playlistListStore := gio.NewListStore(glib.TypeObject)

	pt := &PlaylistsTab{
		Leaflet: builder.GetObject("root").Cast().(*adw.Leaflet),

		currentAdapter: currentAdapter,
		playlistList:   builder.GetObject("playlist-list").Cast().(*gtk.ListView),
	}
	pt.playlistList.SetModel(gtk.NewSingleSelection(playlistListStore))

	pt.ConnectMap(func() { go pt.UpdatePlaylistList() })

	return pt
}

func createPlaylistRow(item *glib.Object) (widget gtk.Widgetter) {
	row := gtk.NewListBoxRow()
	row.SetActionName("app.go-to-playlist")
	// row.SetActionTargetValue(glib.NewVariantString(playlist.ID))
	// label := gtk.NewLabel(playlist.Name)
	label := gtk.NewLabel("Test")
	label.SetCSSClasses([]string{"playlist-list-name"})
	row.SetChild(label)
	row.Show()
	return row
}

// func createPlaylistRow(playlist *base.Playlist) (widget gtk.Widgetter) {
// 	row := gtk.NewListBoxRow()
// 	row.SetActionName("app.go-to-playlist")
// 	row.SetActionTargetValue(glib.NewVariantString(playlist.ID))
// 	label := gtk.NewLabel(playlist.Name)
// 	label.SetCSSClasses([]string{"playlist-list-name"})
// 	row.SetChild(label)
// 	row.Show()
// 	return row
// }

func (pt *PlaylistsTab) UpdatePlaylistList() {
	if pt.currentAdapter == nil {
		return
	}

	playlists, err := pt.currentAdapter.GetPlaylists()
	if err != nil {
		return
	}
	log.Info(playlists)

	// c := pt.playlistList.FirstChild()

	// Add to the playlistListModel

	// for _, playlist := range playlists {
	// 	pt.playlistListModel.Find
	// }

	// END DEBUG
}
