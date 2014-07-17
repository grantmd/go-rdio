package testrdio

import (
	"testing"
)

func TestGetAlbumsByUPC(t *testing.T) {
	c := createClient(t)

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
	c := createClient(t)

	albums, err := c.GetAlbumsForArtist("r49021") // They Might Be Giants
	if err != nil {
		t.Fatal(err)
	}

	if len(albums) == 0 {
		t.Fatal("Album length is 0, but TMBG is very prolific")
	}
}

func TestTestGetAlbumsForLabel(t *testing.T) {
	c := createClient(t)

	_, err := c.GetAlbumsForLabel("l755") // Rhino
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetArtistsForLabel(t *testing.T) {
	c := createClient(t)

	_, err := c.GetArtistsForLabel("l755") // Rhino
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTracksByISRC(t *testing.T) {
	c := createClient(t)

	tracks, err := c.GetTracksByISRC("USPR37300012") // a recording of the song "Love's Theme" by the Love Unlimited Orchestra.
	if err != nil {
		t.Fatal(err)
	}

	if len(tracks) == 0 {
		t.Fatal("Track length is 0, but should be... larger")
	}
}

func TestGetTracksForArtist(t *testing.T) {
	c := createClient(t)

	tracks, err := c.GetTracksForArtist("r49021") // They Might Be Giants
	if err != nil {
		t.Fatal(err)
	}

	if len(tracks) == 0 {
		t.Fatal("Track length is 0, but should be... larger")
	}
}

func TestSearch(t *testing.T) {
	c := createClient(t)

	results, err := c.Search("They Might Be Giants", []string{"Artist"})
	if err != nil {
		t.Fatal(err)
	}

	if len(results) == 0 {
		t.Fatal("Results length is 0, but should be... larger")
	}
}

func TestSearchSuggestions(t *testing.T) {
	c := createClient(t)

	results, err := c.SearchSuggestions("They Might")
	if err != nil {
		t.Fatal(err)
	}

	if len(results) == 0 {
		t.Fatal("Results length is 0, but should be... larger")
	}
}

func TestGetHistoryForUser(t *testing.T) {
	c := createClient(t)

	user, err := c.CurrentUser()
	if err != nil {
		t.Fatal("Could not determine current user")
	}

	sources, err := c.GetHistoryForUser(user.Key, 0, 10)
	if err != nil {
		t.Fatal("Could not get history sources list")
	}

	if len(sources) == 0 {
		t.Fatal("History length probably shouldn't be 0")
	}
}
