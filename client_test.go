package rdio

import (
	"net/url"
	"os"
	"testing"
)

func TestClientNoAuth(t *testing.T) {
	c := &Client{}
	_, err := c.StartAuth()

	if err != ErrDeveloperInactive {
		t.Errorf("Expected 403, got %s\n", err)
	}
}

func TestClientAuth(t *testing.T) {
	c := &Client{
		ConsumerKey:    os.Getenv("RDIO_API_KEY"),
		ConsumerSecret: os.Getenv("RDIO_API_SECRET"),
	}

	if c.ConsumerKey == "" {
		t.Error("Rdio api key is missing (should be in the RDIO_API_KEY environment variable)")
	}

	if c.ConsumerSecret == "" {
		t.Error("Rdio api secret is missing (should be in the RDIO_API_SECRET environment variable)")
	}

	auth, err := c.StartAuth()
	if err != nil {
		t.Fatal(err)
	}

	if auth.Get("login_url") == "" {
		t.Error("login_url missing")
	}

	if auth.Get("oauth_token") == "" {
		t.Error("oauth_token missing")
	}

	if c.Token != auth.Get("oauth_token") {
		t.Errorf("Client and auth token mismatch: %s vs %s", c.Token, auth.Get("oauth_token"))
	}

	if c.TokenSecret == "" {
		t.Error("Client token secret is missing")
	}
}

func TestClientCall(t *testing.T) {
	c := &Client{
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

	params := url.Values{}
	body, err := c.Call("getPlaybackToken", params)
	if err != nil {
		t.Fatal(err)
	}

	s := string(body)

	if s == "" {
		t.Error("Body is empty")
	}
}
