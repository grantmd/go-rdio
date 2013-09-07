package rdio

import (
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
		t.Error(err)
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
