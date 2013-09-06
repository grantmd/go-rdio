package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/grantmd/go-rdio"
	"log"
	"os"
)

func main() {
	// Build a client object
	c := &rdio.Client{}

	// Parse command-line options
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: ./example -consumer_key=foo -consumer_secret=bar\n")
		flag.PrintDefaults()
	}

	flag.StringVar(&c.ConsumerKey, "consumer_key", "", "Your Rdio API consumer key")
	flag.StringVar(&c.ConsumerSecret, "consumer_secret", "", "Your Rdio API consumer secret")
	flag.StringVar(&c.Token, "token", "", "Rdio API user token")
	flag.StringVar(&c.TokenSecret, "token_secret", "", "Rdio API user secret")

	flag.Parse()

	if c.ConsumerKey == "" || c.ConsumerSecret == "" {
		flag.Usage()
		os.Exit(2)
	}

	if c.Token == "" {

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

		fmt.Printf("New token: %s\n", auth.Get("oauth_token"))
		fmt.Printf("New secret: %s\n", auth.Get("oauth_token_secret"))
	}

	// Make our first call
	albums, err := c.GetNewReleases()
	if err != nil {
		log.Fatal(err)
	}

	for _, album := range albums {
		fmt.Printf("%s - %s (http://www.rdio.com%s)\n", album.Name, album.Artist, album.Url)
	}
}
