package cachemiddleware

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"

	"github.com/sumnerevans/sublime-music-next/adapters/base"
)

type CacheMiddleware struct {
	GroundTruthAdapter *base.Adapter

	db *sql.DB
}

func NewCacheMiddleware(dataDir string, groundTruthAdapter *base.Adapter) *CacheMiddleware {
	db, err := sql.Open("sqlite3", dataDir+"/cache.db")
	if err != nil {
		log.Fatal("Could not open cache database.")
	}

	return &CacheMiddleware{
		GroundTruthAdapter: groundTruthAdapter,
		db:                 db,
	}
}

func (cm *CacheMiddleware) Close() {
	cm.db.Close()
}

const getPlaylistsQuery = `
	SELECT id, name, song_count, duration, created, changed, comment, owner, public, cover_art
	FROM playlists
	ORDER BY name
`

func (cm *CacheMiddleware) GetPlaylists() ([]*base.Playlist, error) {
	_, err := cm.db.Query(getPlaylistsQuery)
	if err == nil {
	}

	// Fallback to getting the playlists from the ground truth adapter.
	playlists, err := (*cm.GroundTruthAdapter).GetPlaylists()

	for _, playlist := range playlists {
		log.Info(playlist)
	}

	return playlists, err
}
