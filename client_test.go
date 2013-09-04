package rdio

import (
	"fmt"
	"log"
	"testing"
)

func TestStartAuth(t *testing.T) {
	c := &Client{
		ConsumerKey:    "c9jx7x67bkmqm6ygnjwmbubd",
		ConsumerSecret: "s53DhwNxuq",
	}

	body, err := c.StartAuth()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
}

func TestCompleteAuth(t *testing.T) {
	c := &Client{
		ConsumerKey:    "c9jx7x67bkmqm6ygnjwmbubd",
		ConsumerSecret: "s53DhwNxuq",
	}

	body, err := c.CompleteAuth("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
}

func TestClient(t *testing.T) {
	c := &Client{
		ConsumerKey:    "c9jx7x67bkmqm6ygnjwmbubd",
		ConsumerSecret: "s53DhwNxuq",
	}

	params := make(map[string][]string)
	body, err := c.Call("getNewReleases", params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
}
