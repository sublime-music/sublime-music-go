// Package subsonic defines the Subsonic adapter which implements the Adapter
// API for Subsonic-like backends.
package subsonic

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/sumnerevans/sublime-music-next/adapters/base"
)

type SubsonicAdapter struct {
	Hostname    string
	Username    string
	Password    string
	VerifyCert  bool
	UseSaltAuth bool

	client  *http.Client
	version string

	capabilities *base.Capabilities
}

func New(hostname, username, password string, verifyCert, useSaltAuth bool) (adapter *SubsonicAdapter) {
	adapter = &SubsonicAdapter{
		Hostname:    hostname,
		Username:    username,
		Password:    password,
		VerifyCert:  verifyCert,
		UseSaltAuth: useSaltAuth,

		client:  &http.Client{},
		version: "1.8.0",

		capabilities: &base.Capabilities{
			CanGetPlaylists:       true,
			CanGetPlaylistDetails: true,
			CanCreatePlaylist:     true,
			CanUpdatePlaylist:     true,
			CanDeletePlaylist:     true,
		},
	}
	return
}

// NETWORK PROPERTIES
func (a *SubsonicAdapter) IsNetworked() bool { return true }
func (a *SubsonicAdapter) OnOfflineModeChange(offlineMode bool) {
}

// CAPABILITIES
func (a *SubsonicAdapter) GetCapabilities() *base.Capabilities { return a.capabilities }

// DATA RETRIEVAL METHODS
// Helper methods
func (a *SubsonicAdapter) makeUrl(path string) string {
	return fmt.Sprintf("%s/rest/%s.view", a.Hostname, path)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var numLetters = big.NewInt(int64(len(letters)))

func (a *SubsonicAdapter) randomString(length int) string {
	s := make([]rune, length)
	for i := range s {
		n, _ := rand.Int(rand.Reader, numLetters)
		s[i] = letters[n.Int64()]
	}
	return string(s)
}

func (a *SubsonicAdapter) getParams() map[string]string {
	params := map[string]string{
		"u": a.Username,
		"c": "Sublime Music",
		"f": "json",
		"v": a.version,
	}

	if a.UseSaltAuth {
		// Generates the necessary authentication queryParams to call the
		// Subsonic API. See the Authentication section of
		// www.subsonic.org/pages/api.jsp for more information
		salt := a.randomString(20)
		h := md5.New()
		h.Write([]byte(a.Password + salt))
		params["s"] = salt
		params["t"] = hex.EncodeToString(h.Sum(nil))
	} else {
		params["p"] = a.Password
	}
	return params
}

func (a *SubsonicAdapter) get(url string, timeout *time.Duration, queryParams url.Values) (*http.Response, error) {
	log.Debugf("GET %s %v", url, queryParams)
	// TODO actually use the timeout
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if queryParams == nil {
		queryParams = map[string][]string{}
	}
	for k, v := range a.getParams() {
		queryParams.Add(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()
	return a.client.Do(req)
}

func (a *SubsonicAdapter) getJson(url string, timeout *time.Duration, queryParams url.Values) (*SubsonicResponse, error) {
	resp, err := a.get(url, timeout, queryParams)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if response.SubsonicResponse.Status != "ok" {
		return nil, response.SubsonicResponse.Error
	}

	// Use the version that was sent back in the response for all future
	// requests.
	a.version = response.SubsonicResponse.Version
	return response.SubsonicResponse, nil
}

// Playlists

func (a *SubsonicAdapter) GetPlaylists() ([]*base.Playlist, error) {
	resp, err := a.getJson(a.makeUrl("getPlaylists"), nil, nil)
	if err != nil {
		log.Errorf("Failed to get playlists: %v", err)
		return nil, err
	}

	var playlists []*base.Playlist
	for _, playlist := range resp.Playlists.Playlist {
		playlists = append(playlists, convertPlaylist(playlist))
	}
	sort.Slice(playlists, func(i, j int) bool { return playlists[i].Name < playlists[j].Name })
	return playlists, nil
}

func (a *SubsonicAdapter) GetPlaylistDetails(playlistID string) (*base.Playlist, error) {
	resp, err := a.getJson(a.makeUrl("getPlaylist"), nil, map[string][]string{
		"id": {playlistID},
	})
	if err != nil {
		log.Errorf("Failed to get playlist: %v", err)
		return nil, err
	}

	return convertPlaylist(resp.Playlist), nil
}

func (a *SubsonicAdapter) CreatePlaylist(name string, songIDs []string) (*base.Playlist, error) {
	resp, err := a.getJson(a.makeUrl("createPlaylist"), nil, map[string][]string{
		"name":   {name},
		"sondId": songIDs,
	})
	if err != nil {
		log.Error("Failed to create playlist: %v", err)
		return nil, err
	}
	return convertPlaylist(resp.Playlist), nil
}

func (a *SubsonicAdapter) UpdatePlaylist(playlistID string, name *string, comment *string, public *bool, songIDs []string, appendSongIDs []string) (*base.Playlist, error) {
	if name != nil || comment != nil || public != nil || len(appendSongIDs) > 0 {
		_, err := a.getJson(a.makeUrl("updatePlaylist"), nil, map[string][]string{
			"playlistId":  {playlistID},
			"name":        {*name},
			"comment":     {*comment},
			"public":      {strconv.FormatBool(*public)},
			"songIdToAdd": appendSongIDs,
		})
		if err != nil {
			log.Error("Failed to update playlist: %v", err)
			return nil, err
		}
	}

	var playlist *base.Playlist
	if len(songIDs) > 0 {
		resp, err := a.getJson(a.makeUrl("createPlaylist"), nil, map[string][]string{
			"playlistId": {playlistID},
			"sondId":     songIDs,
		})
		if err != nil {
			log.Error("Failed to update playlist: %v", err)
			return nil, err
		}
		playlist = convertPlaylist(resp.Playlist)
	}

	if playlist == nil {
		return a.GetPlaylistDetails(playlistID)
	}
	return playlist, nil
}

func (a *SubsonicAdapter) DeletePlaylist(playlistID string) error {
	_, err := a.getJson(a.makeUrl("deletePlaylist"), nil, map[string][]string{
		"playlistId": {playlistID},
	})
	if err != nil {
		log.Error("Failed to delete playlist: %v", err)
		return err
	}
	return nil
}
