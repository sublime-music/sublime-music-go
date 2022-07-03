package subsonic

import (
	"fmt"
	"time"
)

type SubsonicIdentifier interface {
	string | int
}

type Song struct {
}

type Playlist struct {
	ID        SubsonicIdentifier         `json:"id"`
	Name      string         `json:"name"`
	SongCount *int           `json:"songCount"`
	Duration  *time.Duration `json:"duration"`
	Songs     []*Song        `json:"entry"`
	Created   *time.Time     `json:"created"`
	Changed   *time.Time     `json:"changed"`
	Comment   *string        `json:"comment"`
	Owner     *string        `json:"owner"`
	Public    *bool          `json:"public"`
	CoverArt  *string        `json:"coverArt"`
}

type Playlists struct {
	Playlist []*Playlist `json:"playlist"`
}

type SubsonicError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *SubsonicError) Error() string {
	return fmt.Sprintf("SubsonicError %d: %s", err.Code, err.Message)
}

func (err *SubsonicError) Is(other error) bool {
	otherErr, ok := other.(*SubsonicError)
	if !ok {
		return false
	}
	return otherErr.Code == err.Code && otherErr.Message == err.Message
}

type SubsonicResponse struct {
	Status string        `json:"status"`
	Error  *SubsonicError `json:"error"`
	Version  string `json:"version"`

	Playlists *Playlists `json:"playlists"`
}

type Response struct {
	SubsonicResponse *SubsonicResponse `json:"subsonic-response"`
}
