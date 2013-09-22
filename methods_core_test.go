package rdio

import (
	"os"
	"testing"
)

func createCoreClient(t *testing.T) (c *Client) {
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

func TestGet(t *testing.T) {
	c := createCoreClient(t)

	objects, err := c.Get([]string{})
	if err != nil {
		t.Fatal(err)
	}

	if len(objects) != 0 {
		t.Errorf("Expected 0 objects, got %d", len(objects))
	}

	objects, err = c.Get([]string{"r49021"})
	if err != nil {
		t.Fatal(err)
	}

	if len(objects) != 1 {
		t.Errorf("Expected 1 object, got %d", len(objects))
	}

	objects, err = c.Get([]string{"r49021", "l755"})
	if err != nil {
		t.Fatal(err)
	}

	if len(objects) != 2 {
		t.Errorf("Expected 2 objects, got %d", len(objects))
	}
}

func TestGetObjectFromShortCode(t *testing.T) {
	c := createCoreClient(t)

	object, err := c.GetObjectFromShortCode("")
	if err == nil {
		t.Error("Expected API error, got none")
	}

	if object != nil {
		t.Errorf("Expected no object, got %s", object)
	}

	object, err = c.GetObjectFromShortCode("QFO0PnCM-A")
	if err != nil {
		t.Fatal(err)
	}

	if object == nil {
		t.Errorf("Expected object, got %s", object)
	}
}

func TestGetObjectFromUrl(t *testing.T) {
	c := createCoreClient(t)

	object, err := c.GetObjectFromUrl("")
	if err == nil {
		t.Error("Expected API error, got none")
	}

	if object != nil {
		t.Errorf("Expected no object, got %s", object)
	}

	object, err = c.GetObjectFromUrl("/artist/Janelle_Mon%C3%A1e/album/The_Electric_Lady/")
	if err != nil {
		t.Fatal(err)
	}

	if object == nil {
		t.Errorf("Expected object, got %s", object)
	}
}
