//
// This is a list of methods that map to Rdio API methods. Calling them returns a proper
// Go struct of the data, which is probably preferable to getting a blob of JSON or
// a byte stream.
//
// Unfortunately, there is a bit of repetition in this code (the Response struct and
// checking the Status == "ok" response), but I don't know how to consolidate it.
//
// List of methods is here: http://www.rdio.com/developers/docs/web-service/methods/
//
// TODO: Some of these methods do not require auth. Do we care?
// TODO: Consistent pagination args for those that do pagination
// TODO: We don't support any of the optional args yet
//

package rdio

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

// Core

// TODO: These are all hard for me, since they can return all different Types, so skipping for now
// Can we use the "type" prop on the results?

// Catalog

func (c *Client) GetAlbumsByUPC(upc string) ([]Album, error) {
	params := url.Values{
		"upc": []string{upc},
	}
	body, err := c.Call("getAlbumsByUPC", params)
	if err != nil {
		return nil, err
	}

	return c.getAlbumResponse(body)
}

func (c *Client) GetAlbumsForArtist(artistKey string) ([]Album, error) {
	params := url.Values{
		"artist": []string{artistKey},
	}
	body, err := c.Call("getAlbumsForArtist", params)
	if err != nil {
		return nil, err
	}

	return c.getAlbumResponse(body)
}

func (c *Client) GetAlbumsForLabel(labelKey string) ([]Album, error) {
	params := url.Values{
		"label": []string{labelKey},
	}
	body, err := c.Call("getAlbumsForLabel", params)
	if err != nil {
		return nil, err
	}

	return c.getAlbumResponse(body)
}

func (c *Client) GetArtistsForLabel(labelKey string) ([]Artist, error) {
	params := url.Values{
		"label": []string{labelKey},
	}
	body, err := c.Call("getArtistsForLabel", params)
	if err != nil {
		return nil, err
	}

	return c.getArtistResponse(body)
}

func (c *Client) GetTracksByISRC(isrc string) ([]Track, error) {
	params := url.Values{
		"isrc": []string{isrc},
	}
	body, err := c.Call("getTracksByISRC", params)
	if err != nil {
		return nil, err
	}

	return c.getTrackResponse(body)
}

func (c *Client) GetTracksForArtist(artistKey string) ([]Track, error) {
	params := url.Values{
		"artist": []string{artistKey},
	}
	body, err := c.Call("getTracksForArtist", params)
	if err != nil {
		return nil, err
	}

	return c.getTrackResponse(body)
}

// TODO: search and searchSuggestions

// Collection

func (c *Client) AddToCollection(keys []string) (bool, error) {
	params := url.Values{
		"keys": []string{strings.Join(keys, ",")},
	}
	body, err := c.Call("addToCollection", params)
	if err != nil {
		return false, err
	}

	return c.getBoolResponse(body)
}

func (c *Client) GetAlbumsForArtistInCollection(artistKey string) ([]Album, error) {
	params := url.Values{
		"artist": []string{artistKey},
	}
	body, err := c.Call("getAlbumsForArtistInCollection", params)
	if err != nil {
		return nil, err
	}

	return c.getAlbumResponse(body)
}

func (c *Client) GetAlbumsInCollection() ([]Album, error) {
	params := url.Values{}
	body, err := c.Call("getAlbumsInCollection", params)
	if err != nil {
		return nil, err
	}

	return c.getAlbumResponse(body)
}

func (c *Client) GetArtistsInCollection() ([]CollectionArtist, error) {
	params := url.Values{}
	body, err := c.Call("getArtistsInCollection", params)
	if err != nil {
		return nil, err
	}

	return c.getCollectionArtistResponse(body)
}

func (c *Client) GetOfflineTracks() ([]Track, error) {
	params := url.Values{}
	body, err := c.Call("getOfflineTracks", params)
	if err != nil {
		return nil, err
	}

	return c.getTrackResponse(body)
}

func (c *Client) GetTracksForAlbumInCollection(albumKey string) ([]Track, error) {
	params := url.Values{
		"album": []string{albumKey},
	}
	body, err := c.Call("getTracksForAlbumInCollection", params)
	if err != nil {
		return nil, err
	}

	return c.getTrackResponse(body)
}

func (c *Client) GetTracksForArtistInCollection(artistKey string) ([]Track, error) {
	params := url.Values{
		"artist": []string{artistKey},
	}
	body, err := c.Call("getTracksForArtistInCollection", params)
	if err != nil {
		return nil, err
	}

	return c.getTrackResponse(body)
}

func (c *Client) GetTracksInCollection() ([]Track, error) {
	params := url.Values{}
	body, err := c.Call("getTracksInCollection", params)
	if err != nil {
		return nil, err
	}

	return c.getTrackResponse(body)
}

func (c *Client) RemoveFromCollection(keys []string) (bool, error) {
	params := url.Values{
		"keys": []string{strings.Join(keys, ",")},
	}
	body, err := c.Call("removeFromCollection", params)
	if err != nil {
		return false, err
	}

	return c.getBoolResponse(body)
}

func (c *Client) SetAvailableOffline(keys []string, offline bool) (bool, error) {
	offlineString := "false"
	if offline {
		offlineString = "true"
	}

	params := url.Values{
		"keys":    []string{strings.Join(keys, ",")},
		"offline": []string{offlineString},
	}
	body, err := c.Call("setAvailableOffline", params)
	if err != nil {
		return false, err
	}

	return c.getBoolResponse(body)
}

// Playlists
// Comments
// Social
// Network
// Activity and Statistics

func (c *Client) GetActivityStream() ([]Album, error) {
	params := url.Values{}
	body, err := c.Call("getActivityStream", params)
	if err != nil {
		return nil, err
	}

	return c.getAlbumResponse(body)
}

func (c *Client) GetHeavyRotationAlbums() ([]Album, error) {
	params := url.Values{
		"type": []string{"albums"},
	}
	body, err := c.Call("getHeavyRotation", params)
	if err != nil {
		return nil, err
	}

	return c.getAlbumResponse(body)
}

func (c *Client) GetHeavyRotationArtists() ([]Artist, error) {
	params := url.Values{
		"type": []string{"artists"},
	}
	body, err := c.Call("getHeavyRotation", params)
	if err != nil {
		return nil, err
	}

	return c.getArtistResponse(body)
}

func (c *Client) GetNewReleases() ([]Album, error) {
	params := url.Values{}
	body, err := c.Call("getNewReleases", params)
	if err != nil {
		return nil, err
	}

	return c.getAlbumResponse(body)
}

func (c *Client) GetTopChartsArtists() ([]Artist, error) {
	params := url.Values{
		"type": []string{"Artist"},
	}
	body, err := c.Call("getTopCharts", params)
	if err != nil {
		return nil, err
	}

	return c.getArtistResponse(body)
}

func (c *Client) GetTopChartsAlbums() ([]Album, error) {
	params := url.Values{
		"type": []string{"Album"},
	}
	body, err := c.Call("getTopCharts", params)
	if err != nil {
		return nil, err
	}

	return c.getAlbumResponse(body)
}

func (c *Client) GetTopChartsTracks() ([]Track, error) {
	params := url.Values{
		"type": []string{"Track"},
	}
	body, err := c.Call("getTopCharts", params)
	if err != nil {
		return nil, err
	}

	return c.getTrackResponse(body)
}

func (c *Client) GetTopChartsPlaylists() ([]Playlist, error) {
	params := url.Values{
		"type": []string{"Playlist"},
	}
	body, err := c.Call("getTopCharts", params)
	if err != nil {
		return nil, err
	}

	return c.getPlaylistResponse(body)
}

// Playback

func (c *Client) GetPlaybackToken() (string, error) {
	params := url.Values{}
	body, err := c.Call("getPlaybackToken", params)
	if err != nil {
		return "", err
	}

	return c.getStringResponse(body)
}

// Private functions for parsing responses

func (c *Client) getPlaylistResponse(body []byte) ([]Playlist, error) {
	type Response struct {
		Status string
		Result []Playlist
	}

	// parse into json
	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Check that we got an OK
	if response.Status != "ok" {
		return nil, errors.New("Got non-ok response from the Rdio API")
	}

	return response.Result, nil
}

func (c *Client) getAlbumResponse(body []byte) ([]Album, error) {
	type Response struct {
		Status string
		Result []Album
	}

	// parse into json
	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Check that we got an OK
	if response.Status != "ok" {
		return nil, errors.New("Got non-ok response from the Rdio API")
	}

	return response.Result, nil
}

func (c *Client) getArtistResponse(body []byte) ([]Artist, error) {
	type Response struct {
		Status string
		Result []Artist
	}

	// parse into json
	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Check that we got an OK
	if response.Status != "ok" {
		return nil, errors.New("Got non-ok response from the Rdio API")
	}

	return response.Result, nil
}

func (c *Client) getTrackResponse(body []byte) ([]Track, error) {
	type Response struct {
		Status string
		Result []Track
	}

	// parse into json
	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Check that we got an OK
	if response.Status != "ok" {
		return nil, errors.New("Got non-ok response from the Rdio API")
	}

	return response.Result, nil
}

func (c *Client) getStringResponse(body []byte) (string, error) {
	type Response struct {
		Status string
		Result string
	}

	// parse into json
	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	// Check that we got an OK
	if response.Status != "ok" {
		return "", errors.New("Got non-ok response from the Rdio API")
	}

	return response.Result, nil
}

func (c *Client) getBoolResponse(body []byte) (bool, error) {
	type Response struct {
		Status string
		Result bool
	}

	// parse into json
	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		return false, err
	}

	// Check that we got an OK
	if response.Status != "ok" {
		return false, errors.New("Got non-ok response from the Rdio API")
	}

	return response.Result, nil
}

func (c *Client) getCollectionArtistResponse(body []byte) ([]CollectionArtist, error) {
	type Response struct {
		Status string
		Result []CollectionArtist
	}

	// parse into json
	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Check that we got an OK
	if response.Status != "ok" {
		return nil, errors.New("Got non-ok response from the Rdio API")
	}

	return response.Result, nil
}
