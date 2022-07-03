// Package adapters provides all of the connectors to the backend services.
package adapters

// The Capabilities struct contains a bunch of booleans that determine what
// things the adapter can be used to do.
//
// These properties can be dynamic based on things such as underlying API
// version, or other factors like that. However, these properties should not be
// dependent on the connection state. After the initial sync, these properties
// are expected to be constant.
type Capabilities struct {
	// Playlists
	// ========================================================================

	// Whether or not the adapter supports getting playlists.
	CanGetPlaylists bool
	// Whether or not the adapter supports getting playlist details.
	CanGetPlaylistDetails bool
	// Whether or not the adapter supports creating a playlist.
	CanCreatePlaylist bool
	// Whether or not the adapter supports updating a playlist.
	CanUpdatePlaylist bool
	// Whether or not the adapter supports deleting a playlist.
	CanDeletePlaylist bool
}

// Adapter is an interface for music backends to implement.
type Adapter interface {
	// Adapter Capabilities
	GetCapabilities() *Capabilities

	// NETWORK PROPERTIES
	// These properties determine whether or not the adapter requires
	// connection over a network and whether the underlying server can be
	// pinged.
	// ========================================================================

	// Whether or not this adapter operates over the network. This will be used
	// to determine whether or not some of the offline/online management
	// features should be enabled.
	IsNetworked() bool

	// This function should be used to handle any operations that need to be
	// performed when Sublime Music goes from online to offline mode or vice
	// versa.
	OnOfflineModeChange(offlineMode bool)

	// DATA RETRIEVAL METHODS
	// These properties determine if what things the adapter can be used to do
	// at the current moment.
	// ========================================================================

	// Playlists

	// Get a list of all of the playlists known by the adapter.
	GetPlaylists() ([]*Playlist, error)

	// Get the details for the given playlistID.
	GetPlaylistDetails(playlistID string) (*Playlist, error)

	// Create a playlist of the given name with the given songs.
	CreatePlaylist(name string, songIDs []string) (*Playlist, error)

	// Update a playlist. If a parameter is nil, then it will be ignored and no
	// updates will occur to that field.
	UpdatePlaylist(playlistID string, name *string, comment *string, public *bool, songIDs []string, appendSongIDs []string) (*Playlist, error)

	// Delete the given playlist.
	DeletePlaylist(playlistID string) error
}
