package testrdio

import (
	"testing"
)

func TestAddToPlaylist(t *testing.T) {
	c := createClient(t)

	_, err := c.AddToPlaylist("", []string{})
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestCreatePlaylist(t *testing.T) {
	c := createClient(t)

	_, err := c.CreatePlaylist("", "", []string{})
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestDeletePlaylist(t *testing.T) {
	c := createClient(t)

	_, err := c.DeletePlaylist("")
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestGetPlaylists(t *testing.T) {
	c := createClient(t)

	_, err := c.GetPlaylists()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserPlaylists(t *testing.T) {
	c := createClient(t)

	_, err := c.GetUserPlaylists("s3318") // Me!
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveFromPlaylist(t *testing.T) {
	c := createClient(t)

	_, err := c.RemoveFromPlaylist("", 0, 1, []string{})
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestSetPlaylistCollaborating(t *testing.T) {
	c := createClient(t)

	_, err := c.SetPlaylistCollaborating("", false)
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestSetPlaylistCollaborationMode(t *testing.T) {
	c := createClient(t)

	_, err := c.SetPlaylistCollaborationMode("", 1)
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestSetPlaylistFields(t *testing.T) {
	c := createClient(t)

	_, err := c.SetPlaylistFields("", "", "")
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}

func TestSetPlaylistOrder(t *testing.T) {
	c := createClient(t)

	_, err := c.SetPlaylistOrder("", []string{})
	if err == nil {
		t.Fatal("Expected API error, got none")
	}
}
