package testrdio

import (
	"testing"
)

func TestGetPlaybackToken(t *testing.T) {
	c := createClient(t)

	token, err := c.GetPlaybackToken()
	if err != nil {
		t.Fatal(err)
	}

	if token == "" {
		t.Error("Token is empty")
	}
}
