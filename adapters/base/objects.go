package base

import (
	"fmt"
	"time"
)

type Song struct {
}

type Playlist struct {
	ID        string
	Name      string
	SongCount *int
	Duration  *time.Duration
	Songs     []*Song
	Created   *time.Time
	Changed   *time.Time
	Comment   *string
	Owner     *string
	Public    *bool
	CoverArt  *string
}

func (p *Playlist) String() string {
	return fmt.Sprintf("Playlist(%s, %s)", p.ID, p.Name)
}
