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
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/sumnerevans/sublime-music-next/adapters"
)

type SubsonicAdapter struct {
	Hostname    string
	Username    string
	Password    string
	VerifyCert  bool
	UseSaltAuth bool

	client  *http.Client
	version string
}

func New(hostname, username, password string, verifyCert, useSaltAuth bool) *SubsonicAdapter {
	return &SubsonicAdapter{
		Hostname:    hostname,
		Username:    username,
		Password:    password,
		VerifyCert:  verifyCert,
		UseSaltAuth: useSaltAuth,

		client:  &http.Client{},
		version: "1.8.0",
	}
}

// NETWORK PROPERTIES
func (a *SubsonicAdapter) IsNetworked() bool { return true }
func (a *SubsonicAdapter) OnOfflineModeChange(offlineMode bool) {
}

// AVAILABILITY PROPERTIES
func (a *SubsonicAdapter) CanGetPlaylists() bool       { return true }
func (a *SubsonicAdapter) CanGetPlaylistDetails() bool { return true }
func (a *SubsonicAdapter) CanCreatePlaylist() bool     { return true }
func (a *SubsonicAdapter) CanUpdatePlaylist() bool     { return true }
func (a *SubsonicAdapter) CanDeletePlaylist() bool     { return true }

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
		// Generates the necessary authentication queryParams to call the Subsonic API
		// See the Authentication section of www.subsonic.org/pages/api.jsp for
		// more information
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

func (a *SubsonicAdapter) get(url string, timeout time.Duration, queryParams url.Values) (*http.Response, error) {
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

func (a *SubsonicAdapter) getJson(url string, timeout time.Duration, queryParams url.Values) (*SubsonicResponse, error) {
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
func (a *SubsonicAdapter) GetPlaylists() ([]*adapters.Playlist, error) {
	if resp, err := a.getJson(a.makeUrl("getPlaylists"), time.Duration(0), nil); err != nil {
		log.Errorf("Failed to get playlists: %v", err)
		return nil, err
	} else {
		var playlists []*adapters.Playlist
		for _, playlist := range resp.Playlists.Playlist {
			playlists = append(playlists, ConvertPlaylist(playlist))
		}
		return playlists, nil
	}
}

func (a *SubsonicAdapter) GetPlaylistDetails(playlistID string) (*adapters.Playlist, error) {
	return nil, nil
}

func (a *SubsonicAdapter) CreatePlaylist(name string, song_ids []string) (*adapters.Playlist, error) {
	return nil, nil
}

func (a *SubsonicAdapter) UpdatePlaylist(playlistID string, name *string, comment *string, public *bool, songIDs []string) (*adapters.Playlist, error) {
	return nil, nil
}

func (a *SubsonicAdapter) DeletePlaylist(playlistID string) error { return nil }
