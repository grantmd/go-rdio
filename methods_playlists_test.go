package rdio

import (
	"os"
	"testing"
)

func createPlaylistClient(t *testing.T) (c *Client) {
	c = &Client{
		ConsumerKey:    os.Getenv("RDIO_API_KEY"),
		ConsumerSecret: os.Getenv("RDIO_API_SECRET"),
		Token:          os.Getenv("RDIO_API_TOKEN"),
		TokenSecret:    os.Getenv("RDIO_API_TOKEN_SECRET"),
	}

	if c.ConsumerKey == "" {
		t.Error("Rdio api key is missing (should be in the RDIO_API_KEY environment variable)")
	}

	if c.ConsumerSecret == "" {
		t.Error("Rdio api secret is missing (should be in the RDIO_API_SECRET environment variable)")
	}

	if c.Token == "" {
		t.Error("Rdio api user token is missing (should be in the RDIO_API_TOKEN environment variable)")
	}

	if c.TokenSecret == "" {
		t.Error("Rdio api user secret is missing (should be in the RDIO_API_TOKEN_SECRET environment variable)")
	}

	return c
}

func TestAddToPlaylist(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.AddToPlaylist("", []string{})
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestCreatePlaylist(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.CreatePlaylist("", "", []string{})
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestDeletePlaylist(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.DeletePlaylist("")
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestGetPlaylists(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.GetPlaylists()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserPlaylists(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.GetUserPlaylists("s3318") // Me!
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveFromPlaylist(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.RemoveFromPlaylist("", 0, 1, []string{})
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestSetPlaylistCollaborating(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.SetPlaylistCollaborating("", false)
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestSetPlaylistCollaborationMode(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.SetPlaylistCollaborationMode("", 1)
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestSetPlaylistFields(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.SetPlaylistFields("", "", "")
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestSetPlaylistOrder(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.SetPlaylistOrder("", []string{})
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}
