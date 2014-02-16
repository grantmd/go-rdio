package testrdio

//
// These are all hard to test, because they operate on the user's collection. So, let's do the minimum for now and hope for the best!
//

import (
	"testing"
)

func TestAddToCollection(t *testing.T) {
	c := createClient(t)

	_, err := c.AddToCollection([]string{""})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAlbumsForArtistInCollection(t *testing.T) {
	c := createClient(t)

	_, err := c.GetAlbumsForArtistInCollection("r49021") // They Might Be Giants
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAlbumsInCollection(t *testing.T) {
	c := createClient(t)

	_, err := c.GetAlbumsInCollection()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetArtistsInCollection(t *testing.T) {
	c := createClient(t)

	_, err := c.GetArtistsInCollection()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetOfflineTracks(t *testing.T) {
	c := createClient(t)

	_, err := c.GetOfflineTracks()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTracksForAlbumInCollection(t *testing.T) {
	c := createClient(t)

	_, err := c.GetTracksForAlbumInCollection("a101334") // Flood, They Might Be Giants
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTracksForArtistInCollection(t *testing.T) {
	c := createClient(t)

	_, err := c.GetTracksForArtistInCollection("r49021") // They Might Be Giants
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTracksInCollection(t *testing.T) {
	c := createClient(t)

	_, err := c.GetTracksInCollection()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveFromCollection(t *testing.T) {
	c := createClient(t)

	_, err := c.RemoveFromCollection([]string{""})
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetAvailableOffline(t *testing.T) {
	c := createClient(t)

	_, err := c.SetAvailableOffline([]string{""}, false)
	if err != nil {
		t.Fatal(err)
	}
}
