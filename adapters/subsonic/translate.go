package subsonic

import (
	"github.com/sumnerevans/sublime-music-next/adapters/base"
)

// Convert a Subsonic API song to an generic Adapter song.
func convertSong(in *Song) (out *base.Song) {
	if in == nil {
		return
	}

	out = &base.Song{}
	return
}

// Convert a Subsonic API playlist to an generic Adapter playlist.
func convertPlaylist(in *Playlist) (out *base.Playlist) {
	if in == nil {
		return
	}

	out = &base.Playlist{}
	out.ID = in.ID.String()
	out.Name = in.Name
	out.SongCount = in.SongCount
	out.Duration = in.Duration.Duration()
	out.Songs = make([]*base.Song, len(in.Songs))
	for i, song := range in.Songs {
		out.Songs[i] = convertSong(song)
	}
	out.Created = in.Created
	out.Changed = in.Changed
	out.Comment = in.Comment
	out.Owner = in.Owner
	out.Public = in.Public
	out.CoverArt = in.CoverArt
	return
}
