package rdio

import (
	"os"
	"testing"
)

func createCatalogClient(t *testing.T) (c *Client) {
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

func TestGetAlbumsByUPC(t *testing.T) {
	c := createCatalogClient(t)

	albums, err := c.GetAlbumsByUPC("011661811324")
	if err != nil {
		t.Fatal(err)
	}

	if len(albums) != 1 {
		t.Fatalf("Album length is %d instead of 1", len(albums))
	}

	if albums[0].Name != "No!" {
		t.Errorf("Album title is %s instead of No!", albums[0].Name)
	}
}

func TestGetAlbumsForArtist(t *testing.T) {
	c := createCatalogClient(t)

	albums, err := c.GetAlbumsByUPC("011661811324")
	if err != nil {
		t.Fatal(err)
	}

	if len(albums) != 1 {
		t.Fatalf("Album length is %d instead of 1", len(albums))
	}

	if albums[0].Name != "No!" {
		t.Errorf("Album title is %s instead of No!", albums[0].Name)
	}
}
