package subsonic

import (
	"github.com/sumnerevans/sublime-music-next/adapters"
)

func ConvertSong(in *Song) (out *adapters.Song) {
	out = &adapters.Song{}
	return
}

func ConvertPlaylist(in *Playlist) (out *adapters.Playlist) {
	out = &adapters.Playlist{}
	out.ID = in.ID.String()
	out.Name = in.Name
	out.SongCount = in.SongCount
	out.Duration = in.Duration
	out.Songs = make([]*adapters.Song, len(in.Songs))
	for i, song := range in.Songs {
		out.Songs[i] = ConvertSong(song)
	}
	out.Created = in.Created
	out.Changed = in.Changed
	out.Comment = in.Comment
	out.Owner = in.Owner
	out.Public = in.Public
	out.CoverArt = in.CoverArt
	return
}
