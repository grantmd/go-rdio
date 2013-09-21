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

	_, err := c.AddToPlaylist("", []string{""})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPlaylists(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.GetPlaylists()
	if err != nil {
		t.Fatal(err)
	}
}

/*func TestGetUserPlaylists(t *testing.T) {
	c := createPlaylistClient(t)

	_, err := c.GetUserPlaylists()
	if err != nil {
		t.Fatal(err)
	}
}*/
