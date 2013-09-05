package main

import (
	"bufio"
	"fmt"
	"github.com/grantmd/go-rdio"
	"log"
	"os"
)

func main() {
	c := &rdio.Client{
		ConsumerKey:    "c9jx7x67bkmqm6ygnjwmbubd",
		ConsumerSecret: "s53DhwNxuq",
	}

	auth, err := c.StartAuth()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Authorize this application at: %s?oauth_token=%s\n", auth.Get("login_url"), auth.Get("oauth_token"))
	fmt.Print("Enter the PIN / OAuth verifier: ")
	bio := bufio.NewReader(os.Stdin)

	verifier, _, err := bio.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	auth, err = c.CompleteAuth(string(verifier))
	if err != nil {
		log.Fatal(err)
	}

	params := make(map[string][]string)
	body, err := c.Call("getNewReleases", params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
}
