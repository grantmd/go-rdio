package main

import (
	"bufio"
	"fmt"
	"github.com/grantmd/go-rdio"
	"log"
	"os"
)

func main() {
	// Build a client object with our keys
	c := &rdio.Client{
		ConsumerKey:    config.ConsumerKey,
		ConsumerSecret: config.ConsumerSecret,
	}

	// Start auth in order to redirect the user to approve our app
	auth, err := c.StartAuth()
	if err != nil {
		log.Fatal(err)
	}

	// Tell the user what to do and wait for their PIN
	fmt.Printf("Authorize this application at: %s?oauth_token=%s\n", auth.Get("login_url"), auth.Get("oauth_token"))
	fmt.Print("Enter the PIN / OAuth verifier: ")
	bio := bufio.NewReader(os.Stdin)

	verifier, _, err := bio.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	// Check their PIN and complete auth so we can make calls
	auth, err = c.CompleteAuth(string(verifier))
	if err != nil {
		log.Fatal(err)
	}

	// Make our first call
	params := make(map[string][]string)
	body, err := c.Call("getNewReleases", params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
}
