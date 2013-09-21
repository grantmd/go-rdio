//
// This is a list of methods that map to Rdio API methods. Calling them returns a proper
// Go struct of the data, which is probably preferable to getting a blob of JSON or
// a byte stream.
//
// Unfortunately, there is a bit of repetition in this code (the Response struct and
// checking the Status == "ok" response), but I don't know how to consolidate it.
//
// List of methods is here: http://developer.rdio.com/docs/read/rest/Methods
//

package rdio

import (
	"encoding/json"
	"errors"
	"net/url"
)

// Core
// Catalog
// Collection
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

	type Response struct {
		Status string
		Result string
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	// Check that we got an OK
	if response.Status != "ok" {
		return "", errors.New("Got non-ok response from the Rdio API")
	}

	return response.Result, nil
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
