package ui

import (
	"io"
	"os"
	"strings"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	log "github.com/sirupsen/logrus"
	"github.com/sumnerevans/sublime-music-next/adapters"
	"github.com/sumnerevans/sublime-music-next/adapters/subsonic"
)

type PlaylistsTab struct {
	*adw.Leaflet

	playlistList *gtk.ListBox
}

func CreatePlaylistsTab() *PlaylistsTab {
	builder := gtk.NewBuilderFromResource("/app/sublimemusic/SublimeMusicNext/ui/playlists.ui")

	playlistScrollWindow := builder.GetObject("playlist-details-content").Cast().(*gtk.ScrolledWindow)
	playlistDetailsFlap := builder.GetObject("playlist-details-flap").Cast().(*adw.Flap)
	vadjustment := playlistScrollWindow.VAdjustment()
	vadjustment.ConnectValueChanged(func() {
		playlistDetailsFlap.SetRevealFlap(vadjustment.Value() > 70)
	})

	pt := &PlaylistsTab{
		Leaflet:      builder.GetObject("root").Cast().(*adw.Leaflet),
		playlistList: builder.GetObject("playlist-list").Cast().(*gtk.ListBox),
	}

	pt.UpdatePlaylistList()

	return pt
}

func createPlaylistRow(playlist *adapters.Playlist) (widget gtk.Widgetter) {
	row := gtk.NewListBoxRow()
	row.SetActionName("app.go-to-playlist")
	row.SetActionTargetValue(glib.NewVariantString(playlist.ID))
	row.SetChild(gtk.NewLabel(playlist.Name))
	row.Show()
	return row
}

func (pt *PlaylistsTab) UpdatePlaylistList() {
	// DEBUG
	passwordFile, _ := os.Open("passwordfile")
	passwordBytes, _ := io.ReadAll(passwordFile)
	password := string(passwordBytes)
	password = strings.TrimSpace(password)

	subsonicAdapter := subsonic.New("https://music.sumnerevans.com", "sumner", password, true, true)

	playlists, err := subsonicAdapter.GetPlaylists()
	if err != nil {
		return
	}
	for _, playlist := range playlists {
		pt.playlistList.Append(createPlaylistRow(playlist))
	}
	log.Info(playlists)
	// for _, playlist := range playlists {
	// 	pt.playlistListModel.Find
	// }

	// END DEBUG
}
