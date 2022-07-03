package subsonic

import (
	"github.com/sumnerevans/sublime-music-next/adapters"
)

func convertSong(in *Song) (out *adapters.Song) {
	out = &adapters.Song{}
	return
}

func convertPlaylist(in *Playlist) (out *adapters.Playlist) {
	if in == nil {
		out = nil
		return
	}

	out = &adapters.Playlist{}
	out.ID = in.ID.String()
	out.Name = in.Name
	out.SongCount = in.SongCount
	out.Duration = in.Duration.Duration()
	out.Songs = make([]*adapters.Song, len(in.Songs))
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
