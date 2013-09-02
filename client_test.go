package rdio

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	c := &Client{
		ConsumerKey:    "c9jx7x67bkmqm6ygnjwmbubd",
		ConsumerSecret: "s53DhwNxuq",
	}

	params := make(map[string][]string)
	body := c.Call("getNewReleases", params)
	fmt.Println(body)
}
