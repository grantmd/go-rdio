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

	albums, err := c.GetAlbumsForArtist("r49021") // They Might Be Giants
	if err != nil {
		t.Fatal(err)
	}

	if len(albums) == 0 {
		t.Fatal("Album length is 0, but TMBG is very prolific")
	}
}

func TestTestGetAlbumsForLabel(t *testing.T) {
	c := createCatalogClient(t)

	_, err := c.GetAlbumsForLabel("l755") // Rhino
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetArtistsForLabel(t *testing.T) {
	c := createCatalogClient(t)

	_, err := c.GetArtistsForLabel("l755") // Rhino
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTracksByISRC(t *testing.T) {
	c := createCatalogClient(t)

	tracks, err := c.GetTracksByISRC("USPR37300012") // a recording of the song "Love's Theme" by the Love Unlimited Orchestra.
	if err != nil {
		t.Fatal(err)
	}

	if len(tracks) == 0 {
		t.Fatal("Track length is 0, but should be... larger")
	}
}

func TestGetTracksForArtist(t *testing.T) {
	c := createCatalogClient(t)

	tracks, err := c.GetTracksForArtist("r49021") // They Might Be Giants
	if err != nil {
		t.Fatal(err)
	}

	if len(tracks) == 0 {
		t.Fatal("Track length is 0, but should be... larger")
	}
}

func TestSearch(t *testing.T) {
	c := createCatalogClient(t)

	results, err := c.Search("They Might Be Giants", []string{"Artist"})
	if err != nil {
		t.Fatal(err)
	}

	if len(results) == 0 {
		t.Fatal("Track length is 0, but should be... larger")
	}
}

func TestSearchSuggestions(t *testing.T) {
	c := createCatalogClient(t)

	results, err := c.SearchSuggestions("They Might")
	if err != nil {
		t.Fatal(err)
	}

	if len(results) == 0 {
		t.Fatal("Track length is 0, but should be... larger")
	}
}

/*
=== RUN TestSearch
method=search&oauth_consumer_key=t5c3whdekw8gtfhr54r45gnn&oauth_nonce=235379&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1379823362&oauth_token=2usw9wfe382y95pfcp5sne9385xahvu4b29bu3dhde5p7uae6qm9qmqrgz2eqpkp&oauth_version=1.0&query=They+Might+Be+Giants&types=Artist
POST&http%3A%2F%2Fapi.rdio.com%2F1%2F&method%3Dsearch%26oauth_consumer_key%3Dt5c3whdekw8gtfhr54r45gnn%26oauth_nonce%3D235379%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1379823362%26oauth_token%3D2usw9wfe382y95pfcp5sne9385xahvu4b29bu3dhde5p7uae6qm9qmqrgz2eqpkp%26oauth_version%3D1.0%26query%3DThey%2BMight%2BBe%2BGiants%26types%3DArtist
--- FAIL: TestSearch (0.06 seconds)
	methods_catalog_test.go:114: 401: Invalid Signature
=== RUN TestSearchSuggestions
method=searchSuggestions&oauth_consumer_key=t5c3whdekw8gtfhr54r45gnn&oauth_nonce=770742&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1379823362&oauth_token=2usw9wfe382y95pfcp5sne9385xahvu4b29bu3dhde5p7uae6qm9qmqrgz2eqpkp&oauth_version=1.0&query=They+Might
POST&http%3A%2F%2Fapi.rdio.com%2F1%2F&method%3DsearchSuggestions%26oauth_consumer_key%3Dt5c3whdekw8gtfhr54r45gnn%26oauth_nonce%3D770742%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1379823362%26oauth_token%3D2usw9wfe382y95pfcp5sne9385xahvu4b29bu3dhde5p7uae6qm9qmqrgz2eqpkp%26oauth_version%3D1.0%26query%3DThey%2BMight
--- FAIL: TestSearchSuggestions (0.04 seconds)
	methods_catalog_test.go:127: 401: Invalid Signature
*/
